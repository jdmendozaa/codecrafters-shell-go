package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		handleError(err)
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		handleError(err)
		command = strings.TrimSpace(command)
		fmt.Printf("%v: command not found\n", command)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
