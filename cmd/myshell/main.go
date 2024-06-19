package main

import (
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtin"
	"log"
	"os"
	"strings"
)

func main() {
	pathEnv := os.Getenv("PATH")
	shellPath := strings.Split(pathEnv, ":")
	BuiltinCommands := builtin.NewBuiltinCommands(shellPath)

	shell := &Shell{
		Path:           shellPath,
		BuiltinCommand: BuiltinCommands,
	}
	shell.Run()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
