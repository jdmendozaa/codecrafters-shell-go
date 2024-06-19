package main

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute(args []string) (string, error)
}

func ExecuteCommand(fullCommand string) {
	commandSplit := strings.Fields(fullCommand)
	command := commandSplit[0]
	args := commandSplit[1:]
	var c Command

	switch command {
	case "exit":
		c = &ExitCommand{}
	default:
		fmt.Fprintf(os.Stderr, "%v: command not found\n", command)
	}
	if c != nil {
		_, err := c.Execute(args)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
	}
}
