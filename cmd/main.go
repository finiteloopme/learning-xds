package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/finiteloopme/goutils/pkg/log"
	client "github.com/finiteloopme/xds-from-scratch/internal/client"
	server "github.com/finiteloopme/xds-from-scratch/internal/server"
	xdsServer "github.com/finiteloopme/xds-from-scratch/pkg/xds/server"
)

func main() {
	appType := flag.String("type", "grpc-server", "grpc-server or grpc-client.  Default is grpc-server")
	host := flag.String("host", "", "Hostname for the gRPC service")
	port := flag.String("port", "", "Port for the gRPC service")
	flag.Parse()
	if *host != "" {
		os.Setenv("GCP_GRPC_HOST", *host)
	}
	if *port != "" {
		os.Setenv("GCP_GRPC_PORT", *port)
	}

	switch *appType {
	case "grpc-server":
		RunServer()
	case "grpc-client":
		RunServer()
	case "xds-server":
		RunXDSServer()
	default:
		log.Fatal(fmt.Errorf("Unable to run the application.  Only valid type is server or client.  \n e.g: xds-from-scratch --type=grpc-server --host=127.0.0.1 --port=9010"))
	}
}

func RunClient() {
	client.RunClient()
}
func RunServer() {
	server.RunServer()
}

func RunXDSServer() {
	xdsServer.DiscoverAndReconcile()
}
