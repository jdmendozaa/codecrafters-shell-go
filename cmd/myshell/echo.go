package main

import (
	"fmt"
	"os"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) Execute(args []string) error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(args, " "))
	if err != nil {
		return err
	}
	return nil
}
