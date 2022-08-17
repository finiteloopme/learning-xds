package main

import (
	"github.com/finiteloopme/goutils/pkg/log"
	server "github.com/finiteloopme/xds-from-scratch/internal/server"
)

func main() {
	log.Info("In main function")
	server.RunServer()
}
