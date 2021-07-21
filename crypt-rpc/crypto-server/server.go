package crypto_server

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Result struct {
	text string
	key  string
}

type Crypto struct {
}

func new() *Crypto {
	c := &Crypto{}
	return c
}

func (c *Crypto) EncryptString(text string, reply *Result) error {

	// generate key
	key := RandStringBytes(22)

	// convert text to bytes
	plainText := []byte(text)

	// process encryption
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}

	nonce := []byte("DistributedSystems")

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)

	reply = &Result{}
	reply.text = fmt.Sprintf("%x", ciphertext)
	reply.key = key

	return nil
}

func (c *Crypto) DecryptString(info *Result, reply *string) error {

	key := info.key
	nonce := []byte("DistributedSystems")
	ciphertext, err := hex.DecodeString(info.text)
	if err != nil {
		panic(err.Error())
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	*reply = string(plaintext)

	return nil
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CryptoServer(address string) {

	// initialize object of rpc server
	cryptoServer := new()
	rpc.Register(cryptoServer)

	// start name server
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Crypto-Server is running")

	for {
		rpc.Accept(l)
	}

}
