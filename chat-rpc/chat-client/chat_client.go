package chat_client

import (
	data "Network-and-Distributed-systems-studies/chat-rpc/chat-server"
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"time"
)

type User struct {
	Name     string
	Message  string
	MsgLidas int
}

func ChatClient(address string, name string) {

	// conecta ao servidor via tcp.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu *string

	// cria o usuário.
	if err := client.Call("Chat.AddUser", name, &resu); err != nil {
		fmt.Printf("Error: in Chat.AddUser %+v", err)
	}

	// função que lida com a input para o servidor.
	go checkData(name, client)

	go senDataToserver(name, client)

	time.Sleep(200000 * time.Second)

}

func senDataToserver(name string, conn *rpc.Client) {

	var resu *string

	data := &data.User{}
	fmt.Println("Digite algo no chat: ")
	fmt.Println("-----------------------------------------------------------")
	for {

		// lendo o teclado
		reader := bufio.NewReader(os.Stdin)
		chat, _ := reader.ReadString('\n')

		// coloca a informação na estrutura de dados aceita pelo servidor
		data.Name = name
		data.Message = fmt.Sprintf("%s:%s", name, chat)

		// envia mensagem para servidor
		if err := conn.Call("Chat.SendMessage", data, &resu); err != nil {
			fmt.Printf("Error: in Chat.SendMessage %+v", err)
			break
		}
	}
}

func checkData(name string, conn *rpc.Client) {
	var resu *bool
	var resu2 *string

	for {

		if err := conn.Call("Chat.NotifyUser", name, &resu); err != nil {
			fmt.Printf("Error: in Chat.NotifyUser %+v\n", err)
			break
		}

		if *resu {

			if err := conn.Call("Chat.ShowMessages", name, &resu2); err != nil {
				fmt.Printf("Error: in Chat.ShowMessages %+v\n", err)
				break
			} else {
				fmt.Println(*resu2)
			}

		}

		time.Sleep(1 * time.Second)
	}

}
