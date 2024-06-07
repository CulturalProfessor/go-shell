package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	colorGreen  = "\033[1;32m"
	colorRed    = "\033[1;31m"
	colorYellow = "\033[1;33m"
	colorReset  = "\033[0m"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(colorGreen + "$ " + colorReset)

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
				fmt.Println(colorRed + "Usage: exit 0" + colorReset)
			}
		case "echo":
			fmt.Println(strings.Join(params, " "))
		case "pwd":
			wd, _ := os.Getwd()
			fmt.Printf("%s\n", wd)
		case "cd":
			if len(params) != 1 {
				fmt.Println(colorRed + "Usage: cd <directory>" + colorReset)
				continue
			}
			changeDirectory(params[0])
		case "type":
			if len(params) != 1 {
				fmt.Println(colorRed + "Usage: type <command>" + colorReset)
				continue
			}
			param := params[0]
			path, ifFound := checkDir(paths, param)
			if param == "echo" || param == "exit" || param == "type" {
				fmt.Printf("%s%s%s is a shell builtin\n", colorGreen, param, colorReset)
			} else {
				if ifFound {
					fmt.Printf("%s%s%s is %s/%s\n", colorGreen, param, colorReset, path, param)
				} else {
					fmt.Printf("%s%s%s: not found\n", colorRed, param, colorReset)
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
					fmt.Printf("%s%s%s: %v\n", colorRed, command, colorReset, err)
				}
			} else {
				fmt.Printf("%s%s%s: command not found\n", colorRed, command, colorReset)
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
	if strings.HasPrefix(path, "~") {
		home := os.Getenv("HOME")
		if home == "" {
			fmt.Println(colorRed + "cd: HOME environment variable is not set" + colorReset)
			return
		}
		if path == "~" {
			path = home
		} else {
			path = filepath.Join(home, path[2:])
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("%scd: %s: Invalid path%s\n", colorRed, path, colorReset)
		return
	}
	err = os.Chdir(absPath)
	if err != nil {
		fmt.Printf("%scd: %s: No such file or directory%s\n", colorRed, path, colorReset)
	}
}
