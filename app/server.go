package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {

		buf := make([]byte, 128)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client: ", err.Error())
				break
			}

			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}
	
		// fmt.Println("read %d bytes", n)
		// fmt.Printf("buf %c", buf[:n])
			
		pongMsg := []byte("+PONG\r\n")
	
		_, err = conn.Write(pongMsg)
		if err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
			os.Exit(1);
		}
			
		// fmt.Println("sent %d bytes", n)
	}

}


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		 
		go handleClient(conn)
	}
	
}
