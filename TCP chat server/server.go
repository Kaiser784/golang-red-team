package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"errors"
)


type server struct{
	rooms map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command),
	}
}

func (server *server) run() {
	for cmd := range server.commands {
		switch cmd.id {
		case cmd_name:
			server.name(cmd.client, cmd.args)
		case cmd_join:
			server.join(cmd.client, cmd.args)
		case cmd_rooms:
			server.listrooms(cmd.client, cmd.args)
		case cmd_msg:
			server.msg(cmd.client, cmd.args)
		case cmd_quit:
			server.quit(cmd.client, cmd.args)
		}
	}
}

func (server *server) newClient(connect net.Conn) {
	log.Printf("New client has connected: %s", connect.RemoteAddr().String())

	c := &client{
		connect: connect,
		name: "anon",
		commands: server.commands,
	}

	c.readInput()
}

func (server *server) name(c *client, args []string) {
	c.name = args[1]
	c.msg(fmt.Sprintf("You are now %s", c.name))
}

func (server *server) join(c *client, args []string) {
	roomname := args[1]

	r, ok := server.rooms[roomname]

	if !ok {
		r = &room{
			name: roomname,
			members: make(map[net.Addr]*client),
		}
		server.rooms[roomname] = r
	}

	r.members[c.connect.RemoteAddr()] = c

	server.leaveroom((c))

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

func (server *server) listrooms(client *client, args []string) {
	var rooms []string
	for name := range server.rooms {
		rooms = append(rooms, name)
	}
	client.msg(fmt.Sprintf("Available rooms are: %s", strings.Join(rooms, "\n")))
}	

func (server *server) msg(client *client, args []string) {
	if client.room == nil {
		client.err(errors.New("Join a room first"))
		return
	}
	client.room.broadcast(client, client.name + " >>> " + strings.Join(args[1:len(args)], " "))
}

func (server *server) quit(client *client, args []string) {
	log.Printf("Client has disconnected: %s", client.connect.RemoteAddr().String())

	server.leaveroom(client)

	client.msg("See You soon")

	client.connect.Close()
}

func (server *server) leaveroom(client *client) {
	if client.room != nil {
		delete(client.room.members, client.connect.RemoteAddr())
		client.room.broadcast(client, fmt.Sprintf("%s has left the room", client.name))
	}
}