package main

import (
	"strings"
	"errors"
)

func executePing() (RespOutput, error) {
	return createRespOutput(PONG, SIMPLESTRING)
}

func executeEcho(args []string) (RespOutput, error) {
	return createRespOutput(args[0], BULKSTRING)
}

func executeSet(args []string) (RespOutput, error) {
	set(args)
	return createRespOutput(OK, SIMPLESTRING)
}

func executeGet(args []string) (RespOutput, error) {
	value := get(args[0])

	if (value == NILSTRING) {
		return createRespOutput(NILSTRING, EMPTY)
	}

	return createRespOutput(value, BULKSTRING)
}

func executeCommand(clientCall ClientCall) (RespOutput, error) {
	switch strings.ToUpper(clientCall.respCommand.command) {
		case PING:
			return executePing()
		case ECHO:
			return executeEcho(clientCall.respCommand.args)
		case SET:
			return executeSet(clientCall.respCommand.args)
		case GET:
			return executeGet(clientCall.respCommand.args)
		default:
			return "", errors.New("Command not implemented yet!")
	}
}