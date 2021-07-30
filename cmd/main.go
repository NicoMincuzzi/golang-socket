package main

func main() {
	client := SocketClient{
		connectionType: "tcp",
		host:           "192.168.1.243",
		port:           "8070",
	}
	client.run()
}
