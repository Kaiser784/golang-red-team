package main

import(
	"fmt"
	"port/port.go"
)

func main(){
	var port int
	var host string
	
	fmt.Print("Enter host you want to scan: ")
	fmt.Scan(&host)

	fmt.Print("Enter Port you want to scan: ")
	fmt.Scan(&port)
	
	fmt.Print()
	fmt.Println("Port Scan Commencing....")
	open := BaseScan("tcp", host, port)

	fmt.Printf("Port Open: %t\n", open)
}