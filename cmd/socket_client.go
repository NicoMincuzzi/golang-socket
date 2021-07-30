package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	message = "Ping\r\n\r\n"
)

type SocketClient struct {
	connectionType string
	host           string
	port           string
}

func New(connectionType string, host string, port string) SocketClient {
	return SocketClient{connectionType, host, port}
}

func (client SocketClient) run() {
	address := client.host + ":" + client.port
	conn, err := net.Dial(client.connectionType, address)
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
