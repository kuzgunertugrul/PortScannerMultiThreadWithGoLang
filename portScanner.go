package main

import(
	"fmt"
	"net"
	"strconv"
	"sync"
)

var openPorts[] string

func portScan(wg *sync.WaitGroup,hostname string, port string) bool {
	defer wg.Done()
	var hostnameWithPort string = hostname + ":" + port

	conn , err := net.Dial("tcp4",hostnameWithPort)
	if err != nil{
		return false
	}else{
		defer fmt.Printf("Host: %s Port: %s Status: Open\n",hostname,port)
		openPorts = append(openPorts,port)
		defer conn.Close()
		return true
	}
}

func main(){
	fmt.Println("Port Scanner")
	
	fmt.Println("Enter target address")
	var hostname string
	fmt.Scanln(&hostname)
	wg := new(sync.WaitGroup)

	wg.Add(65535)
	for i := 1; i<= 65535; i+=15{
		go portScan(wg,hostname,strconv.Itoa(i))
		go portScan(wg,hostname,strconv.Itoa(i+1))
		go portScan(wg,hostname,strconv.Itoa(i+2))
		go portScan(wg,hostname,strconv.Itoa(i+3))
		go portScan(wg,hostname,strconv.Itoa(i+4))
		go portScan(wg,hostname,strconv.Itoa(i+5))
		go portScan(wg,hostname,strconv.Itoa(i+6))
		go portScan(wg,hostname,strconv.Itoa(i+7))
		go portScan(wg,hostname,strconv.Itoa(i+8))
		go portScan(wg,hostname,strconv.Itoa(i+9))
		go portScan(wg,hostname,strconv.Itoa(i+10))
		go portScan(wg,hostname,strconv.Itoa(i+11))
		go portScan(wg,hostname,strconv.Itoa(i+12))
		go portScan(wg,hostname,strconv.Itoa(i+13))
		go portScan(wg,hostname,strconv.Itoa(i+14))
	}
	wg.Wait()
	/* for _, port := range openPorts {
        println(port)
    }	 */
}