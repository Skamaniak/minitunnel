package main

import (
	"flag"
	"github.com/Skamaniak/minitunnel/conf"
	"github.com/Skamaniak/minitunnel/core"
	"github.com/Skamaniak/minitunnel/inbound"
	"github.com/Skamaniak/minitunnel/tunnel"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	configurationFile := flag.String("conf", "./configuration.json", "Configuration file")
	config, err := conf.ReadConfiguration(*configurationFile)
	if err != nil {
		logger.Panicln("Failed to parse configuration", configurationFile, err)
	}

	connections := inbound.CreateInboundListener(config.Listener)
	tunDataStream := tunnel.InitTunInterface(config.Tunnel)

	router := core.NewRouter(connections, tunDataStream)

	router.Route()
}
