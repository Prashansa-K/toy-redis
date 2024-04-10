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
	LF="\n"

	// Commands
	PING = "PING"
	ECHO = "ECHO"
	GET = "GET"
	SET = "SET"
	INFO = "INFO"

	// Requests
	PING_REQUEST = "*1\r\n$4\r\nping\r\n"

	// Command Arguments
	EXPIRYARG = "PX"
	REPLICATION = "REPLICATION"

	// Response Types
	EMPTY = 0
	SIMPLESTRING = 1
	BULKSTRING = 2
	

	// Responses
	EMPTYRESPONSE = "$-1\r\n"
	PONG = "PONG"
	OK = "OK"
	NILSTRING = ""
	ROLE_STRING = "role:"
	MASTER_REPLICATION_ID_STRING = "master_replid:"
	MASTER_REPLICATION_OFFSET_STRING = "master_repl_offset:"
)