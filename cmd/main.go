package main

import (
	//client3 "Network-and-Distributed-systems-studies/calculator-rpc/rpc-client"
	//server3 "Network-and-Distributed-systems-studies/calculator-rpc/rpc-server"
	//client2 "Network-and-Distributed-systems-studies/crypt-rpc/crypto-client"
	//server2 "Network-and-Distributed-systems-studies/crypt-rpc/crypto-server"
	// client1 "Network-and-Distributed-systems-studies/name-rpc/name-client"
	server1 "Network-and-Distributed-systems-studies/name-rpc/name-server"
	// "bufio"
	// "fmt"
	// "os"
	// "strings"
	// "time"
)

func main() {

	// running name server
	server1.NameServer(":8081")

	/*
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Digite o nome do serviço:")
		service, _ := reader.ReadString('\n')
		service = strings.Replace(service, "\n", "", -1)

		fmt.Print("Digite o ip e a porta para o serviço:")
		porta, _ := reader.ReadString('\n')
		porta = strings.Replace(porta, "\n", "", -1)

		client1.RegisterService("127.0.0.1:8081", service, porta)

		// running crypto server
		if service == "crypto" {

			go server2.CryptoServer(porta)

			time.Sleep(2 * time.Second)
		}

		// running calculator server
		if service == "calculator" {

			go server3.RpcServer(porta)

			time.Sleep(2 * time.Second)
		}
	*/
	/*
		fmt.Print("Digite o nome do serviço que deseja utilizar:")
		app, _ := reader.ReadString('\n')
		app = strings.Replace(app, "\n", "", -1)

		server := client1.GetService("127.0.0.1:8081", app)

		if app == "crypto" {

			fmt.Print("Digite o texto a ser encriptado:")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			client2.CryptoClient(server, text)

		}

		if app == "calculator" {

			fmt.Print("Digite o número(N1):")
			var n1 int
			fmt.Scan(&n1)

			fmt.Print("Digite o número(N2):")
			var n2 int
			fmt.Scan(&n2)

			client3.RpcClient(server, n1, n2)
		}

	*/
}
