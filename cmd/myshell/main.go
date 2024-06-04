package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("I/O Error")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]
		params := args[1:]

		switch command {
		case "exit":
			if len(params) == 1 && params[0] == "0" {
				os.Exit(0)
			} else {
				fmt.Println("Usage: exit 0")
			}
		case "echo":
			fmt.Println(strings.Join(params, " "))
		case "pwd":
			wd,_:=os.Getwd()
			fmt.Printf("%s\n",wd)
		case "cd":
			if len(params) != 1 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			changeDirectory(params[0])
		case "type":
			if len(params) != 1 {
				fmt.Println("Usage: type <command>")
				continue
			}
			param := params[0]
			path, ifFound := checkDir(paths, param)
			if param == "echo" || param == "exit" || param == "type" {
				fmt.Printf("%s is a shell builtin\n", param)
			} else {
				if ifFound {
					fmt.Printf("%s is %s/%s\n", param, path, param)
				} else {
					fmt.Printf("%s: not found\n", param)
				}
			}
		default:
			path, ifFound := checkDir(paths, command)
			if ifFound {
				fullPath := filepath.Join(path, command)
				cmd := exec.Command(fullPath, params...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("%s: %v\n", command, err)
				}
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}

func checkDir(paths []string, cmd string) (string, bool) {
	for _, dir := range paths {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.Name() == cmd {
				return dir, true
			}
		}
	}
	return "", false
}

func changeDirectory(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("cd: %s: Invalid path\n", path)
		return
	}
	err = os.Chdir(absPath)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}

