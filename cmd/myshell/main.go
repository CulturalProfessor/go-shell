package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("I/O Error")
		}

		if strings.TrimRight(input,"\n")=="exit 0"{
			os.Exit(0)
			break
		}

		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
	}
}
