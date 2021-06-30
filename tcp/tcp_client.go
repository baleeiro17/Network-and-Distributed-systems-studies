package tcp

import (
	"fmt"
	"net"
)

func TcpClient(connection string, message string) {

	conn, err := net.Dial("tcp", connection)
	if err != nil {
		fmt.Println("error in connect to server")
	}

	n, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error in write in server")
	}

	fmt.Printf("writing %d bytes in tcp client\n", n)

	for {

		buffer := make([]byte, 60)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error in reading of tcp client")
		}

		fmt.Printf("Reading %d bytes in tcp client\n", n)
		fmt.Printf("message received in tcp client is %s\n", buffer[:n])

		// fecha a conexão
		conn.Close()

		return
	}
}
