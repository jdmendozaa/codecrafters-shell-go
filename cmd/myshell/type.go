package main

import (
	"fmt"
	"os"
)

type TypeCommand struct{}

func (c *TypeCommand) Execute(args []string) error {
	command := args[0]

	if _, ok := builtinCommands[command]; ok {
		_, err := fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintf(os.Stderr, "%s: not found\n", command)
		if err != nil {
			return err
		}
	}
	return nil
}
