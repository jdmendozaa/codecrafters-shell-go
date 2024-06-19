package main

import (
	"fmt"
	"os"
	"strings"
)

// Declare here all builtin commands
var builtinCommands = map[string]BuiltinCommand{
	"exit": &ExitCommand{},
	"echo": &EchoCommand{},
	"type": &TypeCommand{},
}

type BuiltinCommand interface {
	Execute(args []string) error
}

func ExecuteBuiltinCommand(fullCommand string) {
	commandSplit := strings.Fields(fullCommand)
	if len(commandSplit) == 0 {
		return
	}
	command := commandSplit[0]
	args := commandSplit[1:]
	var c BuiltinCommand

	if builtinCommand, ok := builtinCommands[command]; ok {
		c = builtinCommand
	} else {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
	}

	if c != nil {
		err := c.Execute(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
