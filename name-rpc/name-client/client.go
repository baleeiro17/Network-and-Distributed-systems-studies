package name_client

import (
	data "Network-and-Distributed-systems-studies/name-rpc/name-server"
	"fmt"
	"net/rpc"
)

func NameClient(address string) {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu *int
	var resu1 data.Service
	args := &data.Service{
		Name: "bank",
		Ip:   "192.168.10.1",
		Port: "2152",
	}

	if err := client.Call("Directory.AddService", args, &resu); err != nil {
		fmt.Printf("Error: in Directory.AddService %+v", err)
	} else {
		fmt.Printf("result of AddService is %d\n", *resu)
	}

	args = &data.Service{
		Name: "oi",
		Ip:   "192.168.10.3",
		Port: "2153",
	}

	if err := client.Call("Directory.AddService", args, &resu); err != nil {
		fmt.Printf("Error: in Directory.AddService %+v", err)
	} else {
		fmt.Printf("result of AddService is %d\n", *resu)
	}

	if err := client.Call("Directory.GetService", "oi", &resu1); err != nil {
		fmt.Printf("Error: in Directory.AddService %+v", err)
	} else {
		fmt.Printf("result of GetService is %s\n", resu1.Ip)
		fmt.Printf("result of GetService is %s\n", resu1.Port)
	}
}
