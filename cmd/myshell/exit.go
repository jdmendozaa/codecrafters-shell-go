package main

import (
	"os"
	"strconv"
)

type ExitCommand struct{}

func (c *ExitCommand) Execute(args []string) error {
	if len(args) == 0 {
		os.Exit(0)
	}
	statusCode := args[0]
	statusCodeInt, err := strconv.Atoi(statusCode)
	if err != nil {
		return err
	}
	os.Exit(statusCodeInt)
	return nil
}
