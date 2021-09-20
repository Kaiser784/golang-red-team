package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	listen = flag.Bool("l", false, "Listen")
	host = flag.String("h", "localhost", "Host")
	port = flag.Int("p", 0, "Port")
)

func main() {
	flag.Parse()

	if *listen {
		startServer()
		return
	}

	if len(flag.Args()) < 2 {
		fmt.Println("Hostname and Port required\n")
		return
	}

	serverHost := flag.Arg(0)
	serverPort := flag.Arg(1)

	startClient(fmt.Sprintf("%s:%s", serverHost, serverPort))

}

func startServer() {

	address := fmt.Sprintf("%s:%d", *host, *port)

	listener, err := net.Listen("tcp", address)
	
	if err != nil {
		panic(err)
	}

	log.Printf("Listening for connections on %s", listener.Addr().String())

	for {
		connect, err := listener.Accept()

		if err != nil {
			log.Printf("Error accepting connection from client: %s\n", err)
		} else {
			go processClient(connect)
		}
	}
}

func startClient(address string) {
	connect, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Printf("Can't connect to serer: %s\n", err)
		return
	}

	_, err = io.Copy(connect, os.Stdin)

	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
}

func processClient(connect net.Conn) {
	_, err := io.Copy(os.Stdout, connect)

	if err != nil {
		fmt.Println(err)
	}
	connect.Close()
}