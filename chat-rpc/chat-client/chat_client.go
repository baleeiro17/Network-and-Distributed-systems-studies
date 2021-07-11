package chat_client

import (
	data "Network-and-Distributed-systems-studies/chat-rpc/chat-server"
	"fmt"
	"net/rpc"
)

type Client struct {
	name    string
	message string
}

func ChatClient(address string) {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu *string

	args := &data.Client{
		Name:    "lucas",
		Message: "Tudo bem?",
	}

	if err := client.Call("Chat.AddUser", "lucas", &resu); err != nil {
		fmt.Printf("Error: in Chat.AddUser %+v", err)
	} else {
		fmt.Printf("result of AddUser is %s\n", *resu)
	}

	if err := client.Call("Chat.AddUser", "Gabriel", &resu); err != nil {
		fmt.Printf("Error: in Chat.AddUser %+v", err)
	} else {
		fmt.Printf("result of AddUser is %s\n", *resu)
	}

	if err := client.Call("Chat.SendMessage", args, &resu); err != nil {
		fmt.Printf("Error: in Chat.SendMessage %+v", err)
	} else {
		fmt.Printf("result of SendMessage is %s\n", *resu)
	}

	if err := client.Call("Chat.ShowMessages", "Gabriel", &resu); err != nil {
		fmt.Printf("Error: in Chat.ShowMessages %+v\n", err)
	} else {
		fmt.Printf("result of ShowMessages is %s\n", *resu)
	}
}
