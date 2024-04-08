package main

import (
	"fmt"
	"strings"
	"errors"
)

type RespCommand struct {
	command string
	args []string
}

type RespOutput string 

const CRLF = "\r\n" // change to backticks to run locally

// Redis client commands are always in array format - an array of bulk strings
// Example:
// "*2\r\n$4\r\necho\r\n$5\r\napple\r\n"
// \r\n or CLRF is the separator used by default in RESP
// * -> indicates array
// 2 -> Length of array
// $ -> bulk string
// 4, 5 -> length of incoming strings respectively (echo, apple)
// first bulk string is always the COMMAND, sometimes the second one is command as well. 
// Subsequent elements of array are args of the command

func parseClientInputResp(input []byte) (RespCommand, error){
	stringInput := string(input)
	tokens := strings.Split(stringInput, CRLF)

	// fmt.Println(tokens);

	respType := tokens[0][0]
	respCommand := RespCommand{}

	if (respType != '*') {
		return respCommand, errors.New("Bulk Array expected!")
	}

	respCommand.command = tokens[2]

	// checking for args
	if len(tokens) > 3 {
		// we need to skip the indices that have length specified
		// index variable gets 0 on first iteration. [3:] slices token from 3rd index
		for index, arg := range tokens[3:] {
			if index % 2 == 1 {
				respCommand.args = append(respCommand.args, arg)
			}
		}
	}

	fmt.Println("cmd:: ", respCommand)

	return respCommand, nil
}


func createRespOutput(output string, respType string) (RespOutput, error){
	respOutput := ""
	switch respType {
		case "simpleString":
			respOutput = fmt.Sprintf("+%s\r\n", output)
		case "bulkString":
			respOutput = fmt.Sprintf("$%d\r\n%s\r\n", len(output), output)
		default:
			return RespOutput(respOutput), errors.New("Not supported")
	}

	return RespOutput(respOutput), nil
}