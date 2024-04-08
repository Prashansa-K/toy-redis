package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

type ClientCall struct {
	conn net.Conn
	respCommand RespCommand
}

func respondClientCall(clientCalls chan ClientCall) {
	for {
		// Wait for a new call in the channel
		newCall := <- clientCalls

		output, err := executeCommand(newCall)
		if err != nil {
			fmt.Println("Error creating output: ", err.Error())
			os.Exit(1);
		}

		// pongMsg := []byte("+PONG\r\n")
	
		_, err = newCall.conn.Write([]byte(output))
		if err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
			os.Exit(1);
		}

		// fmt.Println("sent %d bytes", n)
		// fmt.Printf("read client data %c", newCall.data[0:])
		// fmt.Printf("read from client:: %s ", newCall.data[:len(newCall.data)])
	}
}

func handleClientCall(conn net.Conn, clientCalls chan ClientCall) {
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

		command, err := parseClientInputResp(buf)

		if err != nil {
			fmt.Println("Error occurred in parsing command: ", err.Error())
			break
		}

		// fmt.Println("cmd:: ", command)

		// fmt.Println("read %d bytes", n)
		// fmt.Printf("buf %c", buf[:n])

		// Push client call to channel
		clientCalls <- ClientCall{conn, command}
	}

}

func main() {
	// Initiate server
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	// Create channel
	clientCalls := make(chan ClientCall)
	go respondClientCall(clientCalls)
	

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		 
		go handleClientCall(conn, clientCalls)
	}
	
}
