package socket

import (
	"log"
	"net"
	"os"
)

func New(connectionType string, host string, port string) Client {
	return client{connectionType, host, port}
}

func (c client) Run(conn net.Conn) {
	message := message{
		connection: conn,
	}
	message.messageSender()
	message.messageReceiver()
}

func (c client) Connect() net.Conn {
	address := c.host + ":" + c.port
	conn, err := net.Dial(c.connectionType, address)
	if err != nil {
		log.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	return conn
}
