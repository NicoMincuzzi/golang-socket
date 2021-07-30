package socket

import (
	"fmt"
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
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Cannot send message:", err.Error())
		return
	}
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}
