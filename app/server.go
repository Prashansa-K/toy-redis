package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"flag"
)

var serverConfig struct {
	port string
	address string
	replicaOf string
	role string
}

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

	// os args are fine in case of one argument.
	// For >1, let's use flag so that positioning problem doesn't arise
	// osArgs := os.Args
	// port := DEFAULT_PORT
	// if (len(osArgs) > 1 && osArgs[1] == PORT_ARGUMENT_FLAG) {
	// 	port = osArgs[2]
	// }

	flag.StringVar(&serverConfig.port, "port", DEFAULT_PORT, "port on which redis server would run")
	flag.StringVar(&serverConfig.replicaOf, "replicaof", "", "starts redis server in slave mode")
	flag.Parse() // parses the flags from os.Args

	serverConfig.address = fmt.Sprintf("%s:%s", LOCALHOST, serverConfig.port)

	// No flag passed
	if serverConfig.replicaOf == "" {
		serverConfig.role = MASTER_ROLE
	} else {
		serverConfig.role = SLAVE_ROLE
	}

	listener, err := net.Listen(TCP_PROTOCOL, serverConfig.address)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	fmt.Printf("Server running as %s on port: %s", serverConfig.role, serverConfig.port,)

	defer listener.Close()

	if serverConfig.role == SLAVE_ROLE {
		connectToMaster()
	}

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
