package main

type commandID int

const(
	cmd_name commandID = iota
	cmd_join
	cmd_rooms
	cmd_msg
	cmd_quit
)

type command struct{
	id commandID
	client *client
	args []string
}