package main

import (
	client "Network-and-Distributed-systems-studies/chat-rpc/chat-client"
	server "Network-and-Distributed-systems-studies/chat-rpc/chat-server"
)

func main() {

	// running rpc server
	server.ChatServer(":8081")

	// time.Sleep(5 * time.Second)

	client.ChatClient(":8081", "gabriel")

}
