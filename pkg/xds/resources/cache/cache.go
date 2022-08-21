package cache

import (
	"context"
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cacheV3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	resourceV3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	random "github.com/finiteloopme/goutils/pkg/common"
	log "github.com/finiteloopme/goutils/pkg/log"
)

type XDSCache struct {
	cache           cacheV3.SnapshotCache
	snapshotVersion int
	nodeId          string
	ctx             context.Context
}

func (xdsCache *XDSCache) UpdateSnapshotVersion() int {
	if xdsCache.snapshotVersion == 0 {
		xdsCache.snapshotVersion = random.RandomInt(1, 1000)
	} else {
		xdsCache.snapshotVersion++
	}

	return xdsCache.snapshotVersion
}

func (xdsCache *XDSCache) GetNodeId() string {
	return xdsCache.nodeId
}

func NewCache(_ctx context.Context) *XDSCache {
	var xdsCache XDSCache
	xdsCache.cache = cacheV3.NewSnapshotCache(true, cacheV3.IDHash{}, nil)
	xdsCache.UpdateSnapshotVersion()
	xdsCache.ctx = _ctx
	// xdsCache.nodeId = xdsCache.cache.GetStatusKeys()[0]

	return &xdsCache
}

func (xdsCache *XDSCache) GetCache() cacheV3.SnapshotCache {
	return xdsCache.cache
}

func (xdsCache *XDSCache) UpdateSnapshot(clusters []types.Resource, listeners []types.Resource, secrets []types.Resource, routes []types.Resource, endpoints []types.Resource) {
	version := xdsCache.UpdateSnapshotVersion()
	log.Info("Creating cache with Snapshot Version: " + fmt.Sprint(xdsCache.snapshotVersion))
	var xdsResources map[string][]types.Resource
	if len(xdsResources) == 0 {
		xdsResources = make(map[resource.Type][]types.Resource, 4)
	}
	// clusters
	for _, cluster := range clusters {
		xdsResources[resourceV3.ClusterType] = append(xdsResources[resourceV3.ClusterType], cluster)
	}
	// listeners
	for _, listener := range listeners {
		xdsResources[resourceV3.ListenerType] = append(xdsResources[resourceV3.ListenerType], listener)
	}
	// secrets
	for _, secret := range secrets {
		xdsResources[resourceV3.SecretType] = append(xdsResources[resourceV3.SecretType], secret)
	}
	// routes
	for _, route := range routes {
		xdsResources[resourceV3.RouteType] = append(xdsResources[resourceV3.RouteType], route)
	}
	// endpoint
	for _, endpoint := range endpoints {
		xdsResources[resourceV3.EndpointType] = append(xdsResources[resourceV3.EndpointType], endpoint)
	}
	// xdsResources[resourceV3.ClusterType] = clusters
	// xdsResources[resourceV3.ListenerType] = listeners
	// xdsResources[resourceV3.SecretType] = secrets
	// xdsResources[resourceV3.RouteType] = routes
	// xdsResources[resourceV3.EndpointType] = endpoints

	snap, err := cacheV3.NewSnapshot(fmt.Sprint(version), xdsResources)
	// cant get the consistent snapshot thing working anymore...
	// https://github.com/envoyproxy/go-control-plane/issues/556
	// https://github.com/envoyproxy/go-control-plane/blob/main/pkg/cache/v3/snapshot.go#L110
	// if err := snap.Consistent(); err != nil {
	// 	log.Fatal(fmt.Errorf("snapshot inconsistency: %+v\n%+v", snap, err))
	// 	os.Exit(1)
	// }

	// err = xdsCache.cache.SetSnapshot(context.Background(), xdsCache.nodeId, snap)
	// err = xdsCache.cache.SetSnapshot(xdsCache.ctx, xdsCache.nodeId, snap)
	err = xdsCache.cache.SetSnapshot(xdsCache.ctx, xdsCache.GetCache().GetStatusKeys()[0], snap)
	if err != nil {
		log.Fatal(fmt.Errorf("Could not set snapshot: %v", err))
	}
}

// func (xdsCache *XDSCache) NewSnapshot(nodeId string) {
// 	var xdsResources map[string][]types.Resource
// 	xdsCache.UpdateSnapshotVersion()
// 	snapshot, err := cacheV3.NewSnapshot(
// 		// New version number
// 		fmt.Sprint("%d", xdsCache.snapshotVersion),
// 		xdsResources)
// 	if err != nil {
// 		log.Fatal(fmt.Errorf("Could not create a snapshot: %v", err))
// 	}
// 	log.Info("Creating cache with Snapshot Version: " + fmt.Sprint(xdsCache.snapshotVersion))
// 	xdsCache.cache.SetSnapshot(nodeId, snapshot)
// 	xdsCache.nodeID = nodeId
// }

// func (xdsCache *XDSCache) AddCluster(name string) {

// }

// func CreateSnapshot(c []types.Resource, l []types.Resource, s []types.Resource, nodeId string) {
// 	atomic.AddInt32(&version, 1)
// 	snap := cachev3.NewSnapshot(fmt.Sprint(version), nil, c, nil, l, nil, s)
// 	if err := snap.Consistent(); err != nil {
// 		log.Fatal(fmt.Errorf("snapshot inconsistency: %+v\n%+v", snap, err))
// 	}
// 	err := cache.SetSnapshot(nodeId, snap)
// 	if err != nil {
// 		log.Fatalf("Could not set snapshot %v", err)
// 	}
// }
