package main

import (
	"strings"
	"errors"
)

func executePing() (RespString, error) {
	return createRespString(SIMPLESTRING, PONG)
}

func executeEcho(args []string) (RespString, error) {
	return createRespString(BULKSTRING, args[0])
}

func executeSet(args []string) (RespString, error) {
	set(args)
	return createRespString(SIMPLESTRING, OK)
}

func executeGet(args []string) (RespString, error) {
	value := get(args[0])

	if (value == NILSTRING) {
		return createRespString(EMPTY, NILSTRING)
	}

	return createRespString(BULKSTRING, value)
}

func executeInfo(args []string) (RespString, error) {
	if (strings.ToUpper(args[0]) == REPLICATION) {
		roleString := ROLE_STRING + serverConfig.role

		if (serverConfig.role == SLAVE_ROLE) {
			return createRespString(BULKSTRING, roleString)
		}

		replId := MASTER_REPLICATION_ID_STRING + serverConfig.replicationId
		replOffset := MASTER_REPLICATION_OFFSET_STRING + "0"

		return createRespString(BULKSTRING, roleString, replId, replOffset)
	} else {
		return createRespString(EMPTY, NILSTRING)
	}
}

func executeReplConf(args []string) (RespString, error) {
	return createRespString(SIMPLESTRING, OK)
}

func executeCommand(clientCall ClientCall) (RespString, error) {
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
		case REPLCONF:
			return executeReplConf(clientCall.respCommand.args)
		default:
			return "", errors.New("Command not implemented yet!")
	}
}