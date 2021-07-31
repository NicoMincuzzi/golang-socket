package socket

import (
	"log"
	"net"
)

const (
	defaultMessage = "Ping\r\n\r\n"
)

type message struct {
	connection net.Conn
}

func (m message) messageReceiver() {
	buff := make([]byte, 1024)
	n, _ := m.connection.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}

func (m message) messageSender() {
	_, err := m.connection.Write([]byte(defaultMessage))
	if err != nil {
		log.Println("Cannot send defaultMessage:", err.Error())
	}
	log.Printf("Send: %s", defaultMessage)
}
