package main

import (
	client "Network-and-Distributed-systems-studies/name-rpc/name-client"
	server "Network-and-Distributed-systems-studies/name-rpc/name-server"
	"time"
)

func main() {

	// running calculator-rpc server
	go server.NameServer(":8081")

	time.Sleep(1 * time.Second)

	client.NameClient("127.0.0.1:8081")
}
