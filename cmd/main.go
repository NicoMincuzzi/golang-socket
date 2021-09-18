package main

import "golang-socket/cmd/socket"

func main() {
	client := socket.New("tcp", "192.168.1.244", "8070")

	conn := client.Connect()
	defer conn.Close()
	client.Run(conn)
}
