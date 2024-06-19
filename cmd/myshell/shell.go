package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtin"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Shell struct {
	Path           []string
	BuiltinCommand *builtin.Commands
}

func (shell *Shell) Run() {
	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		handleError(err)
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		handleError(err)
		commandSplit := strings.Fields(command)
		if len(commandSplit) == 0 {
			continue
		}
		commandName := commandSplit[0]
		args := commandSplit[1:]

		// We need to handle the type command here, since it needs access to the Shell available commands
		if commandName == "type" {
			shell.typeCommand(args[0])
			continue
		}

		// Check if command is a builtin
		if _, ok := shell.BuiltinCommand.CommandsMap[commandName]; ok {
			err := shell.BuiltinCommand.ExecuteBuiltinCommand(commandName, args...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			continue
		}

		// Check if command is in path
		commandPath, err := shell.SearchInPath(commandName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		shell.RunExternalProgram(context.Background(), commandPath, args)
	}
}

func (shell *Shell) typeCommand(command string) {

	if _, ok := shell.BuiltinCommand.CommandsMap[command]; ok {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
	} else if fullPath, err := shell.SearchInPath(command); err == nil {
		fmt.Fprintf(os.Stdout, "%s is %s\n", command, fullPath)
	} else {
		fmt.Fprintf(os.Stderr, "%s: not found\n", command)
	}
}

func (shell *Shell) SearchInPath(command string) (string, error) {

	for _, shellPath := range shell.Path {
		fullPath := path.Join(shellPath, command)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", CommandNotFoundError{command}
}

func (shell *Shell) RunExternalProgram(ctx context.Context, commandPath string, args []string) {
	output, err := exec.CommandContext(ctx, commandPath, args...).CombinedOutput()
	fmt.Fprint(os.Stdout, string(output))

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
