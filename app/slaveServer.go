package main

import (
	"fmt"
	"net"
	"os"
	"flag"
)

var masterConfig struct {
	host string
	port string
}

func connectToMaster() {
	masterConfig.host = serverConfig.replicaOf
	masterConfig.port = flag.Args()[len(flag.Args()) - 1] // last argument

	connection, err := net.Dial(TCP_PROTOCOL, masterConfig.host + ":" + masterConfig.port)
	if (err != nil) {
		fmt.Println("Error connecting to master: ", err.Error())
		os.Exit(1)
	}
	// defer connection.Close()

	// starting handshake
	pingMaster(connection)

	// 1st replconf
	sendReplConfToMaster(connection, REPLCONF, "listening-port", serverConfig.port)
	// 2nd replconf
	sendReplConfToMaster(connection, REPLCONF, "capa", "psync2")

	sendPsyncToMaster(connection)
}

func pingMaster (connection net.Conn) {
	pingRequest, _ := createRespString(RESPARRAY, PING)

	_, err := connection.Write([]byte(pingRequest))
	if err != nil {
		fmt.Println("Error sending ping to master: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 128)
	_, err = connection.Read(buf)
	if err != nil {
		fmt.Println("Received error from master: ", err.Error())
		os.Exit(1)
	}
}

func sendReplConfToMaster (connection net.Conn, replConfCommand ...string) {
	replConfRequest, _ := createRespString(RESPARRAY, replConfCommand...)

	_, err := connection.Write([]byte(replConfRequest))
	if err != nil {
		fmt.Println("Error sending replconf to master: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 128)
	_, err = connection.Read(buf)
	if err != nil {
		fmt.Println("Received error from master: ", err.Error())
		os.Exit(1)
	}
}

func sendPsyncToMaster (connection net.Conn) {
	psyncRequest, _ := createRespString(RESPARRAY, PSYNC, "?", "-1")

	_, err := connection.Write([]byte(psyncRequest))
	if err != nil {
		fmt.Println("Error sending replconf to master: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 128)
	_, err = connection.Read(buf)
	if err != nil {
		fmt.Println("Received error from master: ", err.Error())
		os.Exit(1)
	}
}