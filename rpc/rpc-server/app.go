package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"
)

func RpcServer(address string) {

	// initialize handler and register in rpc server.
	handler := new()
	rpc.Register(handler)

	// start rpc server
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Server is running")

	for {
		rpc.Accept(l)
	}

}
