package main

import (
	"strings"
	"errors"
)

func executePing() (RespOutput, error) {
	return createRespOutput(SIMPLESTRING, PONG)
}

func executeEcho(args []string) (RespOutput, error) {
	return createRespOutput(BULKSTRING, args[0])
}

func executeSet(args []string) (RespOutput, error) {
	set(args)
	return createRespOutput(SIMPLESTRING, OK)
}

func executeGet(args []string) (RespOutput, error) {
	value := get(args[0])

	if (value == NILSTRING) {
		return createRespOutput(EMPTY, NILSTRING)
	}

	return createRespOutput(BULKSTRING, value)
}

func executeInfo(args []string) (RespOutput, error) {
	if (strings.ToUpper(args[0]) == REPLICATION) {
		roleString := ROLE_STRING + serverConfig.role

		if (serverConfig.role == SLAVE_ROLE) {
			return createRespOutput(BULKSTRING, roleString)
		}
		
		replId := MASTER_REPLICATION_ID_STRING + "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
		replOffset := MASTER_REPLICATION_OFFSET_STRING + "0"

		return createRespOutput(BULKSTRING, roleString, replId, replOffset)
	} else {
		return createRespOutput(EMPTY, NILSTRING)
	}
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
		case INFO:
			return executeInfo(clientCall.respCommand.args)
		default:
			return "", errors.New("Command not implemented yet!")
	}
}