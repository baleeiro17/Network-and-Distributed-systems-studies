package main

import (
	client "Network-and-Distributed-systems-studies/rpc/rpc-client"
	server "Network-and-Distributed-systems-studies/rpc/rpc-server"
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	if len(args) <= 2 {
		fmt.Println("Please provide enough arguments.")
		return
	}

	fmt.Printf("N1:%s\n", args[1])
	fmt.Printf("N2:%s\n", args[2])

	n1, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	n2, err := strconv.Atoi(args[2])
	if err != nil {
		return
	}

	// running rpc server
	server.RpcServer(":8081")

	// time.Sleep(5 * time.Second)

	// running rpc client
	client.RpcClient("127.0.0.1:8081", n1, n2)
}
