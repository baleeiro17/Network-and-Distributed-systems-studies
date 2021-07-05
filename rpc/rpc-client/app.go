package rpc_client

import (
	data "Network-and-Distributed-systems-studies/rpc/rpc-server"
	"fmt"
	"net/rpc"
)

func RpcClient(address string, n1 int, n2 int) {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu *int
	args := &data.Valor{
		N1: n1,
		N2: n2,
	}

	if err := client.Call("Handler.Sum", args, &resu); err != nil {
		fmt.Printf("Error: in Handler.Sum %+v", err)
	} else {
		fmt.Printf("result of sum is %d\n", *resu)
	}

	if err := client.Call("Handler.Subtraction", args, &resu); err != nil {
		fmt.Printf("Error: in Handler.Sum %+v", err)
	} else {
		fmt.Printf("result of subtraction is %d\n", *resu)
	}

	if err := client.Call("Handler.Division", args, &resu); err != nil {
		fmt.Printf("Error: in Handler.Sum %+v", err)
	} else {
		fmt.Printf("result of division is %d\n", *resu)
	}

	if err := client.Call("Handler.Product", args, &resu); err != nil {
		fmt.Printf("Error: in Handler.Sum %+v", err)
	} else {
		fmt.Printf("result of product is %d\n", *resu)
	}
}
