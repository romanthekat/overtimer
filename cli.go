package main

import (
	"fmt"
	"os"
)

type commandLineArg string

const (
	start  commandLineArg = "start"
	stop   commandLineArg = "stop"
	spend  commandLineArg = "spend"
	status commandLineArg = "status"
)

func readCommand() (commandLineArg, error) {
	return parseArguments(os.Args[1:])
}

func parseArguments(args []string) (commandLineArg, error) {
	if len(args) > 1 {
		return status, fmt.Errorf("only one parameter can be specified")
	} else if len(args) == 0 {
		return status, nil
	}

	command := commandLineArg(args[0])
	switch command {
	case start, stop, spend, status:
		return command, nil
	case "":
		return status, nil
	default:
		return status, fmt.Errorf("unknown command provided")
	}
}
