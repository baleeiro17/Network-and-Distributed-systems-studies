package main

import (
	client "Network-and-System-Distributed-studies/rpc/rpc-client"
	server "Network-and-System-Distributed-studies/rpc/rpc-server"
)

func main() {

	// running rpc server
	server.RpcServer("127.0.0.1:8081")

	// time.Sleep(5 * time.Second)

	// running rpc client
	client.RpcClient("127.0.0.1:8081", 4, 2)

}
