package main

import (
	"flag"
	"os"

	"github.com/finiteloopme/goutils/pkg/log"
	client "github.com/finiteloopme/xds-from-scratch/internal/client"
)

func main() {
	host := flag.String("host", "", "Hostname for the gRPC service")
	port := flag.String("port", "", "Port for the gRPC service")
	flag.Parse()
	if *host != "" {
		os.Setenv("GCP_GRPC_HOST", *host)
	}
	if *port != "" {
		os.Setenv("GCP_GRPC_PORT", *port)
	}
	log.Info("In main function")
	client.RunClient()
}
