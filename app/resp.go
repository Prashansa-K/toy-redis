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

type RespString string

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


func createRespString(respType int, output ...string) (RespString, error){
	respOutput := ""
	switch respType {
		case SIMPLESTRING:
			for _, outputString := range output {
				respOutput += fmt.Sprintf("+%s%s", outputString, CRLF)
			}

		case BULKSTRING:
			if len(output) == 1 {
				respOutput = fmt.Sprintf("$%d%s%s%s", len(output[0]), CRLF, output[0], CRLF)
			} else {
				totalRespStringLength := 0
				
				for _, outputString := range output {
					totalRespStringLength += len(outputString)
					respOutput += fmt.Sprintf("%s%s", outputString, LF)
				}

				// Adding length of LFs added in between each output string
				totalRespStringLength = totalRespStringLength + len(output) 

				// Adding length of total string in front
				respOutput  = fmt.Sprintf("$%d%s%s%s", totalRespStringLength, CRLF, respOutput, CRLF)
			}

		case RESPARRAY:
			totalLenOfRespArray := len(output)

			for _, outputString := range output {
				respOutput += fmt.Sprintf("$%d%s%s%s", len(outputString), CRLF, outputString, CRLF)
			}

			// Adding length of array in front
			respOutput = fmt.Sprintf("*%d%s%s", totalLenOfRespArray, CRLF, respOutput)

		case EMPTY:
			respOutput = EMPTYRESPONSE
			
		default:
			return RespString(respOutput), errors.New("Not supported")
	}

	return RespString(respOutput), nil
}