// Stark-Proxy: Sprint: 1, Tiket: 1.1 TCP Listener Initialization

package main

import (
	"fmt"
	"net"
)

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
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		_ = n
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
	}
}
