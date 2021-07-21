package name_server

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Service struct {
	Name string
	Ip   string
	Port string
}

type Directory struct {
	services       sync.Map // map[name]*Service, name as key
	authentication sync.Map // map[name]*Token, name as key
}

func new() *Directory {
	d := &Directory{}
	return d
}

func (d *Directory) AddService(service *Service, reply *int) error {

	fmt.Println("Added service to server:")
	fmt.Println("Service Name:", service.Name)
	fmt.Println("Service IP:", service.Ip)
	fmt.Println("Service Port:", service.Port)

	// persist the service
	d.services.Store(service.Name, service)

	// generate token.
	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(10000001)

	// persist the token
	d.authentication.Store(service.Name, token)

	fmt.Println("Sending token to access server")
	*reply = token

	return nil
}

func (d *Directory) GetService(service string, reply *Service) error {

	// get the service information
	fmt.Println("Search to service to server:")
	app, ok := d.services.Load(service)
	if !ok {
		fmt.Println("Service was not found")
		reply = nil
	}

	// return the service
	reply.Name = app.(*Service).Name
	reply.Ip = app.(*Service).Ip
	reply.Port = app.(*Service).Port

	return nil
}

func NameServer(address string) {

	// initialize object of rpc server
	nameServer := new()
	rpc.Register(nameServer)

	// start name server
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Name-Server is running")

	for {
		rpc.Accept(l)
	}

}
