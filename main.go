package main

import (
	"flag"
	"fmt"
)

func main() {

	//keep main minimal

	//flags
	serverPort := flag.String("port", ":8080", "specify the port the server listens on")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	tls := flag.Bool("tls", true, "run HTTPS")
	flag.Parse()

	server, err := NewServer(*serverPort)
	if err != nil {
		panic(err)
	}

	if !*tls   {
		server.Start()
	} else {
		server.StartTLS(*certFile, *keyFile)
	}
	fmt.Println("ended")

}
