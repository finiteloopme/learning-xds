package server

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	clusterV3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointV3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerV3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routeV3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	routerV3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcmV3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	clusterSvc "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discoverySvc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	endpointSvc "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	listenerSvc "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routeSvc "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	secretSvc "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	serverV3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	log "github.com/finiteloopme/goutils/pkg/log"
	xCache "github.com/finiteloopme/xds-from-scratch/pkg/xds/resources/cache"
	ads "github.com/finiteloopme/xds-from-scratch/pkg/xds/server/ads"
	"github.com/golang/protobuf/ptypes"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	// log "github.com/sirupsen/logrus"
)

type XDSServer struct {
	Server               serverV3.Server
	AggregatesDS         ads.Callbacks
	Hostname             string `default:"127.0.0.1"`
	Port                 int32  `default:"18000"`
	Ctx                  context.Context
	MaxConcurrentStreams uint32 `default:"1000000"`
	XCache               xCache.XDSCache
	// Ads (aggregated) mode ?
	AdsMode bool `default:"true"`
}

func NewXDSServer(cb *ads.Callbacks) *XDSServer {
	// signal := make(chan struct{})
	// cb := &ads.Callbacks{
	// 	Signal:   signal,
	// 	Fetches:  0,
	// 	Requests: 0,
	// }
	ctx := context.Background()
	cache := *xCache.NewCache(ctx)

	xdsServer := XDSServer{
		Server: serverV3.NewServer(
			ctx,
			cache.GetCache(),
			cb,
		),
		AggregatesDS: ads.Callbacks{},
		Ctx:          ctx,
		XCache:       cache,
	}
	envconfig.Process("XDS", &xdsServer)
	return &xdsServer
}

func (server *XDSServer) GetURI() string {
	return fmt.Sprint(server.Hostname, ":", server.Port)
}

const grpcMaxConcurrentStreams = 1000

func (server *XDSServer) RunXdsServer() {
	opts := []grpc.ServerOption{grpc.MaxConcurrentStreams(server.MaxConcurrentStreams)}
	// if creds, err := xdscreds.NewServerCredentials(
	// 	xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()}); err != nil {
	// }
	grpcServer := grpc.NewServer(opts...)
	listener, err := net.Listen("tcp", fmt.Sprint(server.GetURI()))
	if err != nil {
		log.Fatal(fmt.Errorf("Error listening on [%v]: %v", server.GetURI(), err))
	}
	if server.AdsMode {
		// Register ADS
		discoverySvc.RegisterAggregatedDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered ADS")
	} else {
		// Register individual services
		endpointSvc.RegisterEndpointDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered eDS")
		clusterSvc.RegisterClusterDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered cDS")
		routeSvc.RegisterRouteDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered rDS")
		listenerSvc.RegisterListenerDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered lDS")
		secretSvc.RegisterSecretDiscoveryServiceServer(grpcServer, server.Server)
		log.Info("Registered sDS")
	}
	log.Info("xDS Management Server started at: " + server.GetURI())

	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-server.Ctx.Done()
	grpcServer.GracefulStop()
}

const (
	backendHostName = "be.cluster.local"
	listenerName    = "be-srv"
	routeConfigName = "be-srv-route"
	clusterName     = "be-srv-cluster"
	virtualHostName = "be-srv-vs"
)

// UpstreamPorts is a type that implements flag.Value interface
type UpstreamPorts []int

// String is a method that implements the flag.Value interface
func (u *UpstreamPorts) String() string {
	// See: https://stackoverflow.com/a/37533144/609290
	return strings.Join(strings.Fields(fmt.Sprint(*u)), ",")
}

// Set is a method that implements the flag.Value interface
func (u *UpstreamPorts) Set(port string) error {
	log.Info("[UpstreamPorts] " + port)
	i, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	*u = append(*u, i)
	return nil
}

var upstreamPorts UpstreamPorts = UpstreamPorts{50051, 50052, 50053}

func DiscoverAndReconcile() {
	signal := make(chan struct{})
	cb := &ads.Callbacks{
		Signal:   &signal,
		Fetches:  0,
		Requests: 0,
	}
	server := NewXDSServer(cb)
	go server.RunXdsServer()
	<-signal
	server.AggregatesDS.Report()

	var lbendpoints []*endpointV3.LbEndpoint
	currentHost := 0
	nodeId := server.XCache.GetCache().GetStatusKeys()[0]
	log.Info(">>>>>>>>>>>>>>>>>>> creating NodeID: " + nodeId)
	for {
		if currentHost+1 <= len(upstreamPorts) {

			v := upstreamPorts[currentHost]
			currentHost++
			// ENDPOINT
			log.Info(">>>>>>>>>>>>>>>>>>> creating ENDPOINT [remoteHost:port] " + backendHostName + ":" + fmt.Sprint(v))
			hst := &core.Address{Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Address:  backendHostName,
					Protocol: core.SocketAddress_TCP,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: uint32(v),
					},
				},
			}}

			epp := &endpointV3.LbEndpoint{
				HostIdentifier: &endpointV3.LbEndpoint_Endpoint{
					Endpoint: &endpointV3.Endpoint{
						Address: hst,
					}},
				HealthStatus: core.HealthStatus_HEALTHY,
			}
			lbendpoints = append(lbendpoints, epp)

			endpoints := []types.Resource{
				&endpointV3.ClusterLoadAssignment{
					ClusterName: clusterName,
					Endpoints: []*endpointV3.LocalityLbEndpoints{{
						Locality: &core.Locality{
							Region: "us-central1",
							Zone:   "us-central1-a",
						},
						Priority:            0,
						LoadBalancingWeight: &wrapperspb.UInt32Value{Value: uint32(1000)},
						LbEndpoints:         lbendpoints,
					}},
				},
			}

			// CLUSTER
			log.Info(">>>>>>>>>>>>>>>>>>> creating CLUSTER " + clusterName)
			clusters := []types.Resource{
				&clusterV3.Cluster{
					Name:                 clusterName,
					LbPolicy:             clusterV3.Cluster_ROUND_ROBIN,
					ClusterDiscoveryType: &clusterV3.Cluster_Type{Type: clusterV3.Cluster_EDS},
					EdsClusterConfig: &clusterV3.Cluster_EdsClusterConfig{
						EdsConfig: &core.ConfigSource{
							ConfigSourceSpecifier: &core.ConfigSource_Ads{},
						},
					},
				},
			}

			// RDS
			log.Info(">>>>>>>>>>>>>>>>>>> creating RDS " + virtualHostName)

			routes := []types.Resource{
				&routeV3.RouteConfiguration{
					Name:             routeConfigName,
					ValidateClusters: &wrapperspb.BoolValue{Value: true},
					VirtualHosts: []*routeV3.VirtualHost{{
						Name:    virtualHostName,
						Domains: []string{listenerName}, //******************* >> must match what is specified at xds:/// //
						Routes: []*routeV3.Route{{
							Match: &routeV3.RouteMatch{
								PathSpecifier: &routeV3.RouteMatch_Prefix{
									Prefix: "",
								},
							},
							Action: &routeV3.Route_Route{
								Route: &routeV3.RouteAction{
									ClusterSpecifier: &routeV3.RouteAction_Cluster{
										Cluster: clusterName,
									},
								},
							},
						},
						},
					}},
				},
			}

			// LISTENER
			log.Info(">>>>>>>>>>>>>>>>>>> creating LISTENER " + listenerName)
			hcRds := &hcmV3.HttpConnectionManager_Rds{
				Rds: &hcmV3.Rds{
					RouteConfigName: routeConfigName,
					ConfigSource: &core.ConfigSource{
						ResourceApiVersion: core.ApiVersion_V3,
						ConfigSourceSpecifier: &core.ConfigSource_Ads{
							Ads: &core.AggregatedConfigSource{},
						},
					},
				},
			}

			hff := &routerV3.Router{}
			tctx, err := ptypes.MarshalAny(hff)
			if err != nil {
				log.Fatal(fmt.Errorf("could not unmarshall router: %v\n", err))
			}

			manager := &hcmV3.HttpConnectionManager{
				CodecType:      hcmV3.HttpConnectionManager_AUTO,
				RouteSpecifier: hcRds,
				HttpFilters: []*hcmV3.HttpFilter{{
					Name: wellknown.Router,
					ConfigType: &hcmV3.HttpFilter_TypedConfig{
						TypedConfig: tctx,
					},
				}},
			}

			pbst, err := ptypes.MarshalAny(manager)
			if err != nil {
				panic(err)
			}

			listeners := []types.Resource{
				&listenerV3.Listener{
					Name: listenerName,
					ApiListener: &listenerv3.ApiListener{
						ApiListener: pbst,
					},
				}}

			server.XCache.UpdateSnapshot(clusters, listeners, nil, routes, endpoints)
			// Sleep and check for updates
			time.Sleep(10 * time.Second)
		}
	}
}
