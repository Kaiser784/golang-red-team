package port

import(
	"net"
	"time"
	"strconv"
)

type ScanResult struct{
	port int
	state string
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

func BaseScan(hostname string) []ScanResult {
	
	var results []ScanResult

	for i:=0; i<=1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}