package socket

import (
	"log"
	"net"
)

func messageReceiver(conn net.Conn) {
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}

func messageSender(conn net.Conn) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("Cannot send message:", err.Error())
	}
	log.Printf("Send: %s", message)
}
