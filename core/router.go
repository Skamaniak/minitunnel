package core

import (
	"github.com/sirupsen/logrus"
	"net"
)

var logger = logrus.New()

type Router struct {
	inboundConnections chan net.Conn
	tunDataStream      chan []byte
}

func (r *Router) Route() {
	logger.Info("Starting routing routine")

	for connection := range r.inboundConnections {
		//TODO
		logger.Info(connection)
	}
}

func NewRouter(inboundConnections chan net.Conn, tunDataStream chan []byte) Router {
	return Router{
		inboundConnections: inboundConnections,
		tunDataStream:      tunDataStream,
	}
}
