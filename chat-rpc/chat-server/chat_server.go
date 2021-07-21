package chat_server

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type User struct {
	Name     string
	Message  string
	MsgLidas int
}

type Chat struct {
	idMsg    int
	users    sync.Map // map[name]*User, name como key
	messages sync.Map // map[int]*String, idMsg como Key
}

// chat allocation.
func new() *Chat {
	h := &Chat{}
	h.idMsg = 0
	return h
}

func (c *Chat) AddUser(name string, reply *string) error {

	// cria o usuário e cadastra no chat.
	user := &User{}
	user.Name = name
	user.MsgLidas = 0

	// grava o usuário na gerência do chat.
	c.users.Store(name, user)

	// insere a mensagem no banco de dados do chat.
	msg := fmt.Sprintf("%s: %s\n", name, "Olá a todos, cheguei!!!")
	c.messages.Store(c.idMsg, msg)

	// incrementa o identificador da mensagem.
	c.idMsg++

	*reply = "usuário cadastrado com sucesso!"

	return nil
}

func (c *Chat) SendMessage(client User, reply *string) error {

	// verifica o usuário.
	user, ok := c.users.Load(client.Name)
	if !ok {
		*reply = "usuário não existe"
	}

	// atualiza os bancos de todos os usuários com a mensagem.
	msg := client.Message

	// insere a mensagem no banco de dados do chat.
	c.messages.Store(c.idMsg, msg)

	// incrementa o identificador da mensagem do chat.
	c.idMsg++

	//incrementa a quantidade de mensagens lidas do usuário que envia.
	user.(*User).MsgLidas++

	result := "mensagem enviada com sucesso!"
	*reply = result

	return nil
}

func (c *Chat) ShowReadMessages(client string, reply *int) error {

	// verifica o usuário.
	user, ok := c.users.Load(client)
	if !ok {
		*reply = 0
	}

	*reply = user.(*User).MsgLidas

	return nil
}

func (c *Chat) NotifyUser(client string, reply *bool) error {

	// verifica o usuário.
	user, ok := c.users.Load(client)
	if !ok {
		return fmt.Errorf("invalid user")
	}

	// retorna para o cliente um true se uma nova mensagem foi enviada.
	if user.(*User).MsgLidas != c.idMsg {
		*reply = true
	} else {
		*reply = false
	}

	return nil
}

func (c *Chat) ShowMessages(client string, reply *string) error {

	// verifica o usuário.
	user, ok := c.users.Load(client)
	if !ok {
		*reply = "usuário não existe"
	}

	if user.(*User).MsgLidas != c.idMsg {
		*reply = ""
		for i := user.(*User).MsgLidas; i < c.idMsg; i++ {
			msg, ok := c.messages.Load(i)
			if ok {
				*reply = fmt.Sprintf("%s\n%s", *reply, msg)
			}
		}

		// atualiza as mgs lidas.
		user.(*User).MsgLidas = c.idMsg

	} else {
		*reply = ""
	}

	return nil
}

func ChatServer(address string) {

	// initialize handler and register in calculator-rpc server.
	chat := new()
	rpc.Register(chat)

	// start calculator-rpc server
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Server is running")

	for {
		rpc.Accept(l)
	}

}
