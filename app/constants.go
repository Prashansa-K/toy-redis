package main

const (
	TCP_PROTOCOL = "tcp"
	LOCALHOST = "0.0.0.0"
	DEFAULT_PORT = "6379"
	PORT_ARGUMENT_FLAG = "--port"

	CRLF = "\r\n" // change to backticks to run locally

	// Commands
	PING = "PING"
	ECHO = "ECHO"
	GET = "GET"
	SET = "SET"

	// Command Arguments
	EXPIRYARG = "PX"

	// Response Types
	SIMPLESTRING = "simpleString"
	BULKSTRING = "bulkString"
	EMPTY = "empty"

	// Responses
	EMPTYRESPONSE = "$-1\r\n"
	PONG = "PONG"
	OK = "OK"
	NILSTRING = ""
)