package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		handleError(err)
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		handleError(err)
		ExecuteCommand(command)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
