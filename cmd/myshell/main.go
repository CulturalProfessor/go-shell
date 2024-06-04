package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("I/O Error")
		}
		_, param, _ := strings.Cut(input, " ")
		param = strings.TrimRight(param, "\n")
		switch strings.TrimRight(input, "\n") {
		case "exit 0":
			os.Exit(0)
			return
		case "echo " + param:
			fmt.Println(param)
		case "type "+param:
			path, ifFound := checkDir(paths, param)
			if ifFound {
				fmt.Printf("%s is %s\n", param, path)
			} else {
				fmt.Printf("%s: %s\n", param, path)
			}
		default:
			fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
		}

	}
}

func checkDir(paths []string, cmd string) (string, bool) {
	path, ifFound := "command not found", false
	for i := 0; i < len(paths); i++ {
		entries, _ := os.ReadDir(paths[i])
		for _, e := range entries {
			if e.Name() == cmd {
				path, ifFound = paths[i], true
				return path, ifFound
				} else {
				path, ifFound = "command not found", false
			}
		}
	}
	return path, ifFound
}
