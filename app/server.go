package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	pongMsg := []byte("+PONG\r\n")

	n, err := conn.Write(pongMsg)
	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		os.Exit(1);
	}
	
	fmt.Println("sent %d bytes", n)
	
}
