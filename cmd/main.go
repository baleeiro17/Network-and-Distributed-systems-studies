package main

import (
	client "Network-and-Distributed-systems-studies/chat-rpc/chat-client"
	server "Network-and-Distributed-systems-studies/chat-rpc/chat-server"
)

func main() {

	// running rpc server
	server.ChatServer(":8081")

	// time.Sleep(5 * time.Second)

	client.ChatClient("127.0.0.1:8081", "marcos")

}
