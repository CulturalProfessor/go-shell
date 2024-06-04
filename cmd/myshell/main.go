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
		_,param,_:=strings.Cut(input," ")
		param=strings.TrimRight(param,"\n")
		
		switch strings.TrimRight(input,"\n") {
		case "exit 0":
			os.Exit(0)
			return
		case "echo "+param:
			fmt.Println(param)
		default:
			fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
		}

	}
}
