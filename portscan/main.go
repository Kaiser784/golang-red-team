package main

import(
	"fmt"
	"github.com/Kaiser784/golang-red-team/portscan/port"
)

func main(){
	// var p int
	var host string
	
	fmt.Print("Enter host you want to scan: ")
	fmt.Scan(&host)

	// fmt.Print("Enter Port you want to scan: ")
	// fmt.Scan(&p)
	
	fmt.Println("Port Scan Commencing....")
	// open := 
	port.BaseScan(host)

	// fmt.Println(open)
}