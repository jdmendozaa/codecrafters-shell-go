package main

import "fmt"

type CommandNotFoundError struct {
	Command string
}

func (e CommandNotFoundError) Error() string {
	return fmt.Sprintf("%s: command not found", e.Command)
}
