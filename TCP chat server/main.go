package main

import (
	"log"
	"net"
)

func main() {
	server := newServer()
	go server.run()

	listener, err := net.Listen("tcp", ":8833")
	
	if err != nil {
		log.Fatalf("unable to start server : %s", err.Error())
	}

	defer listener.Close()

	log.Printf("started server on : 8833")

	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go server.newClient(connect)
	}
}