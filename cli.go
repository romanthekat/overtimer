package main

import (
	"fmt"
	"os"
)

type commandLineArg string

const (
	start  commandLineArg = "start"
	stop   commandLineArg = "stop"
	status commandLineArg = "status"
)

func readCommandLineArgOs() (commandLineArg, error) {
	return readCommandLineArg(os.Args[1:])
}

func readCommandLineArg(args []string) (commandLineArg, error) {
	if len(args) > 1 {
		return status, fmt.Errorf("only one parameter can be specified")
	} else if len(args) == 0 {
		return status, nil
	}

	command := commandLineArg(args[0])
	switch command {
	case start:
		return start, nil
	case stop:
		return stop, nil
	case status, "":
		return status, nil
	default:
		return status, fmt.Errorf("unknown command provided")
	}
}
