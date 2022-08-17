package main

import (
	"github.com/finiteloopme/goutils/pkg/log"
	client "github.com/finiteloopme/xds-from-scratch/internal/client"
)

func main() {
	log.Info("In main function")
	client.RunClient()
}
