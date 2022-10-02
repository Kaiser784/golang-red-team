package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func execute(command string, args []string) (err error) {
	arg := args
	cmd := exec.Command(command, arg...)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

func main() {
	// works only with single argument commands like "ls" 
	// doesn't work for "ls -al"
	cmd := flag.String("cmd","","Select Command to use")
	flag.Parse()
	command := "sudo"
	execute(command, []string{*cmd})
}