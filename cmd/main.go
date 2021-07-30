package main

import "golang-socket/cmd/socket"

func main() {
	client := socket.New("tcp", "192.168.1.243", "8070")
	client.Run()
}
