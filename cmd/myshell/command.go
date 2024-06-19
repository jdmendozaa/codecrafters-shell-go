package main

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute(args []string) error
}

func ExecuteCommand(fullCommand string) {
	commandSplit := strings.Fields(fullCommand)
	if len(commandSplit) == 0 {
		return
	}
	command := commandSplit[0]
	args := commandSplit[1:]
	var c Command

	switch command {
	case "exit":
		c = &ExitCommand{}
	case "echo":
		c = &EchoCommand{}
	default:
		fmt.Fprintf(os.Stderr, "%v: command not found\n", command)
	}

	if c != nil {
		err := c.Execute(args)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
	}
}
