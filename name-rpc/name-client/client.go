package name_client

import (
	data "Network-and-Distributed-systems-studies/name-rpc/name-server"
	"fmt"
	"net/rpc"
)

func RegisterService(address string, name string, ip string) {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu *int

	args := &data.Service{
		Name: name,
		Ip:   ip,
	}

	if err := client.Call("Directory.AddService", args, &resu); err != nil {
		fmt.Printf("Error: in Directory.AddService %+v", err)
	} else {
		fmt.Printf("Token of the service is %d\n", *resu)
	}

}

func GetService(address string, name string) string {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu1 data.Service

	if err := client.Call("Directory.GetService", name, &resu1); err != nil {
		fmt.Printf("Error: in Directory.AddService %+v", err)
	} else {
		return resu1.Ip
	}

	return "not found"
}
