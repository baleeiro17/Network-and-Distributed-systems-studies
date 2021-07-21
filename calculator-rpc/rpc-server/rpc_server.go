package rpc_server

import (
	"fmt"
	"net"
	"net/rpc"
)

// data structures used in RPC server.
type Valor struct {
	N1 int
	N2 int
}

type Handler struct {
}

// Handler allocation.
func new() *Handler {
	h := &Handler{}
	return h
}

// remote procedures function used in RPC server.
func (*Handler) Sum(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 + values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Subtraction(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 - values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Division(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 / values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Product(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 * values.N2

	// return the value
	*reply = resul

	return nil
}

func RpcServer(address string) {

	// initialize handler and register in calculator-rpc server.
	handler := new()
	rpc.Register(handler)

	// start calculator-rpc server
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Calculator-Server is running")

	for {
		rpc.Accept(l)
	}

}
