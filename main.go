package main

import (
	"flag"
	"fmt"
)

func main() {

	//keep main minimal

	//flags
	serverPort := flag.String("port", ":8080", "specify the port the server listens on")
	certFile := flag.String("certfile", "cmd/self-signed-cert/cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "cmd/self-signed-cert/key.pem", "key PEM file")
	flag.Parse()

	server, err := NewServer(*serverPort)
	if err != nil {
		panic(err)
	}
	//server.Start()

	server.StartTLS(*certFile, *keyFile)

	fmt.Println("ended")

}
