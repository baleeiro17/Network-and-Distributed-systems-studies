package chat_server

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type Client struct {
	Name    string
	Message string
}

type User struct {
	name     string
	idMsg    int
	messages sync.Map // map[int]*String idMsg as Key
}

type Chat struct {
	name     [30]string
	idUsr    int
	database sync.Map // map[name]*User, name as key
}

// chat allocation.
func new() *Chat {
	h := &Chat{}
	h.idUsr = 0
	return h
}

func (c *Chat) AddUser(name string, reply *string) error {

	// cria o usuário e cadastra no chat.
	user := &User{}
	user.name = name
	user.idMsg = 0

	// grava o usuário na gerência do chat.
	// e também os nomes que são usados como chave.
	c.database.Store(name, user)
	c.name[c.idUsr] = name
	c.idUsr++

	// atualiza os bancos de todos os usuários com a mensagem de boa vindas do usuário.
	msg := fmt.Sprintf("%s: %s", name, "Olá a todos, cheguei!!!")
	for i := 0; i < c.idUsr; i++ {
		user, ok := c.database.Load(c.name[i])
		if ok {
			idMsg := user.(*User).idMsg
			user.(*User).messages.Store(idMsg, msg)
			user.(*User).idMsg++
		}
	}

	*reply = name

	return nil
}

func (c *Chat) SendMessage(message Client, reply *string) error {

	// atualiza os bancos de todos os usuários com a mensagem.
	msg := fmt.Sprintf("%s: %s", message.Name, message.Message)

	// primeiro checar todos os clientes que existem.
	for i := 0; i < c.idUsr; i++ {
		user, ok := c.database.Load(c.name[i])
		if ok {
			idMsg := user.(*User).idMsg
			user.(*User).messages.Store(idMsg, msg)
			user.(*User).idMsg++
		}
	}

	result := "user found"
	*reply = result

	return nil
}

func (c *Chat) ShowMessages(client string, reply *string) error {

	user, ok := c.database.Load(client)
	if ok {
		for x := 0; x < user.(*User).idMsg; x++ {
			message, find := user.(*User).messages.Load(x)
			if find {
				fmt.Println(fmt.Sprintf("%s", message))
			}
		}
	} else {
		*reply = "user not found"
	}

	return nil
}

func ChatServer(address string) {

	// initialize handler and register in rpc server.
	chat := new()
	rpc.Register(chat)

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
