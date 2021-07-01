package main

import (
	"Network-and-System-Distributed-studies/tcp"
	"fmt"
	"os"
	"time"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	fmt.Println("Server running in", args[1])

	// server TCP.
	go tcp.TcpServer(args[1])

	time.Sleep(5 * time.Second)

	// client TCP that connects to server.
	tcp.TcpClient(args[1], "SD é muito interessante!")

}
