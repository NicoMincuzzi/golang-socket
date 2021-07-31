package socket

import (
	"log"
	"net"
	"os"
)

const (
	message = "Ping\r\n\r\n"
)

type client struct {
	connectionType string
	host           string
	port           string
}

func New(connectionType string, host string, port string) client {
	return client{connectionType, host, port}
}

func (c client) Run() {
	address := c.host + ":" + c.port
	conn, err := net.Dial(c.connectionType, address)
	if err != nil {
		log.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	messageSender(conn)
	messageReceiver(conn)
}
