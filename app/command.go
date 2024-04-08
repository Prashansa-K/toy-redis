package main

import (
	"strings"
	"errors"
)

func executePing() (RespOutput, error) {
	return createRespOutput("PONG", "simpleString")
}

func executeEcho(args []string) (RespOutput, error) {
	return createRespOutput(args[0], "bulkString")
}

func executeCommand(clientCall ClientCall) (RespOutput, error) {
	switch strings.ToLower(clientCall.respCommand.command) {
		case "ping":
			return executePing()
		case "echo":
			return executeEcho(clientCall.respCommand.args)
		default:
			return "", errors.New("Command not implemented yet!")
	}
}