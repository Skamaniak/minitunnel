package inbound

import (
	"fmt"
	"github.com/Skamaniak/minitunnel/conf"
	"github.com/sirupsen/logrus"
	"net"
)

var logger = logrus.New()

func acceptConnections(listener net.Listener, connections chan net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			close(connections)
			logger.Panicln("Failed to accept inbound connection", err)
		}
		logger.Info("Accepted new connection from", conn.RemoteAddr())
		connections <- conn
	}
}

func CreateInboundListener(config conf.ListenerConfiguration) chan net.Conn {
	connectionType := config.ConnectionType
	address := fmt.Sprintf("%v:%v", config.Host, config.Port)
	logger.Infof("Creating inbound %v listener on %v", connectionType, address)

	listener, err := net.Listen(connectionType, address)
	if err != nil {
		logger.Panicln("Failed to create inbound listener", err)
	}

	connections := make(chan net.Conn)
	go acceptConnections(listener, connections)
	return connections
}
