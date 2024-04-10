package main

const (
	// Defaults
	TCP_PROTOCOL = "tcp"
	LOCALHOST = "0.0.0.0"
	DEFAULT_PORT = "6379"
	MASTER_ROLE = "master"
	SLAVE_ROLE = "slave"

	// OS ARGS
	PORT_ARGUMENT_FLAG = "--port"
	PORT_ARGUMENT_FLAG_SHORT = "-p"

	CRLF = "\r\n" // change to backticks to run locally

	// Commands
	PING = "PING"
	ECHO = "ECHO"
	GET = "GET"
	SET = "SET"
	INFO = "INFO"

	// Command Arguments
	EXPIRYARG = "PX"
	REPLICATION = "REPLICATION"

	// Response Types
	SIMPLESTRING = "simpleString"
	BULKSTRING = "bulkString"
	EMPTY = "empty"

	// Responses
	EMPTYRESPONSE = "$-1\r\n"
	PONG = "PONG"
	OK = "OK"
	NILSTRING = ""
	ROLE_STRING = "role:"
)