package tunnel

import (
	"github.com/Skamaniak/minitunnel/conf"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func readPackets(dataStream chan []byte) {
	//TODO
}

func InitTunInterface(config conf.TunnelConfiguration) chan []byte {
	logger.Info("Creating system TUN Interface with", config.IpAddresses)

	dataStream := make(chan []byte)
	go readPackets(dataStream)
	return dataStream
}
