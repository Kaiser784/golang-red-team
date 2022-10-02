package port

import(
	"fmt"
	"net"
	"time"
	"strconv"
)

type ScanResult struct{
	port 	int
	state 	string
}

func ScanPort(protocol, hostname string, port int) ScanResult{
	
	result:= ScanResult{port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn,err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil{
		result.state = "Closed"
		return result
	}
	defer conn.Close()
	result.state = "Open"
	return result
}

func BaseScan(hostname string) {
	
	var results []ScanResult
	var flag int = 0

	for i:=0; i<=1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))

		if results[i].state == "Open" {
			fmt.Print("Port - ", results[i].port)
			fmt.Printf(" : %s\n", results[i].state)
			flag++
		}
	}
	if flag == 0 {
		fmt.Println("All ports are closed or filtered")
	}

	// return results
}