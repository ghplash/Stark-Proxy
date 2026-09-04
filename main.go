// Stark-Proxy. Sprint: 1, Tiket: 1.1 TCP Listener Initialization ||AND|| Tiket: 1.2 HTTP protocol parser ||AND|| Tiket: 1.3

package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"sync"
)

type CachePoll struct {
	key    string
	valye  []byte
	muSyns sync.RWMutex
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error in server:", err)
		return
	}
	fmt.Println("Proxy server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	request, err := http.ReadRequest(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("URl:", request.URL, "Header:", request.Header)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	resp.Write(conn)
}

// Stark-Proxy. Sprint: 2, Tiket: 2.1 Thread-safe LRU memory manager
