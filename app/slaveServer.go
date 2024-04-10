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

	connection.Write([]byte(PING_REQUEST))
}