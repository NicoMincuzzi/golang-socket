package socket

import (
	"log"
	"net"
	"os"
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
	conn := c.connect()

	defer conn.Close()

	message := message{
		connection: conn,
	}
	message.messageSender()
	message.messageReceiver()
}

func (c client) connect() net.Conn {
	address := c.host + ":" + c.port
	conn, err := net.Dial(c.connectionType, address)
	if err != nil {
		log.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	return conn
}
