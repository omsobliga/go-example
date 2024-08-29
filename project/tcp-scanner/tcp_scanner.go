// go run tcp_scanner.go --host baidu.com --startPort 0 --endPort 100
package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func TcpScanner(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Println("scan", address)
	conn, err := net.DialTimeout("tcp", address, time.Microsecond * 200)
	if err == nil {
		conn.Close()
		return true
	}
	return true
}

func SyncScan() {
	parts := []int{}
	for port := 0; port < 100; port++ {
		if TcpScanner("baidu.com", port) {
			parts = append(parts, port)
		}
	}
	fmt.Println(parts)
}

func AsyncScan(hostname string, startPort, endPort int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	parts := []int{}
	for port := startPort; port < endPort; port++ {
		wg.Add(1)
		go func(port int) {
			res := TcpScanner(hostname, port)
			if res {
				mutex.Lock()
				parts = append(parts, port)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Println(parts)
}

func main() {
	hostname := flag.String("host", "", "host")
	startPort := flag.Int("startPort", 0, "start port")
	endPort := flag.Int("endPort", 0, "end port")
	flag.Parse()

	fmt.Println("AsyncScan start")
	AsyncScan(*hostname, *startPort, *endPort)
}
