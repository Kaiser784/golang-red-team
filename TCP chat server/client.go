package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct{
	connect		net.Conn
	name 		string
	room		*room
	commands chan<- command
}

func (client *client) readInput() {
	for {
		msg, err := bufio.NewReader(client.connect).ReadString('\n')

		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd{
		case "/name": 
			client.commands <- command{
				id: cmd_name,
				client: client,
				args: args,
			}
		case "/join":
			client.commands <- command{
				id: cmd_join,
				client: client,
				args: args,
			}
		case "/rooms":
			client.commands <- command{
				id: cmd_rooms,
				client: client,
				args: args,
			}
		case "/msg":
			client.commands <- command{
				id: cmd_msg,
				client: client,
				args: args,
			}
		case "/quit":
			client.commands <- command{
				id: cmd_quit,
				client: client,
				args: args,
			}

		default:
			client.err(fmt.Errorf("unknown command: %s", cmd))

		}
	}
}

func (client *client) err(err error) {
	client.connect.Write([]byte("ERR " + "err.Error" + "\n"))
}

func (client *client) msg(msg string) {
	client.connect.Write([]byte(">>> " + msg + "\n"))
}