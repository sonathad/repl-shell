package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		checkCmd(cmd)
	}
}

func checkCmd(cmd string) {
	builtins := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
	}

	line := strings.Split(cmd, " ")
	suf := strings.Join(line[1:], " ")

	switch pre := line[0]; pre {
	case "exit":
		os.Exit(0)
	case "":
		return
	case "echo":
		fmt.Println(strings.TrimSpace(suf))
	case "type":
		if builtins[suf] {
			fmt.Printf("%s is a shell builtin\n", suf)
		} else {
			fmt.Printf("%s: not found\n", suf)
		}
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
}
