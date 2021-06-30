package tcp

import (
	"fmt"
	"net"
)

func TcpServer(connection string) {

	ln, err := net.Listen("tcp", connection)
	if err != nil {
		// handle error
		fmt.Println("error in listen of TCP server")
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	for {

		buffer := make([]byte, 60)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error in reading of tcp server")
		}

		fmt.Printf("Reading %d bytes in tcp server\n", n)

		fmt.Printf("Receiving %s message\n", buffer[:n])

		// send the response to the client.
		msg := fmt.Sprintf("%s%s%s", "Mensagem ", buffer[:n], " recebida com sucesso!")

		n, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("error in writing of tcp server")
		}
	}
}
