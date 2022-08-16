package main

import (
	server "github.com/finiteloopme/xds-from-scratch/internal"
	"github.com/finiteloopme/goutils/pkg/log"
)

func main() {
	log.Info("In main function")
	server.RunServer()
}
