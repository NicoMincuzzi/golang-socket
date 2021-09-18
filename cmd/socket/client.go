package socket

import "net"

type client struct {
	connectionType string
	host           string
	port           string
}

type Client interface {
	Connect() net.Conn
	Run(connection net.Conn)
}
