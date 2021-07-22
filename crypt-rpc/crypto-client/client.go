package crypto_client

import (
	data "Network-and-Distributed-systems-studies/crypt-rpc/crypto-server"
	"fmt"
	"net/rpc"
)

func CryptoClient(address string, text string) {

	// connect to server.
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	var resu1 data.Result
	var resu *string

	if err := client.Call("Crypto.EncryptString", text, &resu1); err != nil {
		fmt.Printf("Error: in Crypto.EncryptString %+v", err)
	} else {
		fmt.Printf("Chave para desencriptar o texto: %s\n", resu1.Key)
		fmt.Printf("Texto encriptado: %s\n", resu1.Text)
	}

	if err := client.Call("Crypto.DecryptString", resu1, &resu); err != nil {
		fmt.Printf("Error: in Crypto.DecryptString %+v", err)
	} else {
		fmt.Printf("Frase desencriptada novamente: %s\n", *resu)
	}

}
