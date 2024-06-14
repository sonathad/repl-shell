package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
			cmdPath, err := exec.LookPath(suf)
			if err != nil {
				fmt.Printf("%s: not found\n", suf)
			} else {
				fmt.Printf("%s is %s\n", suf, cmdPath)
			}
		}
	case "cd":
		if strings.Contains(suf, "~") {
			homedir, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}
			suf = strings.Replace(suf, "~", homedir, -1)
		}
		err := os.Chdir(suf)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", suf)
		}
	default:
		handleNativeCommands(line)
	}
}

func handleNativeCommands(line []string) {
	var cmdRes *exec.Cmd
	if len(line) > 1 {
		cmdRes = exec.Command(line[0], strings.Join(line[1:], " "))
	} else {
		cmdRes = exec.Command(line[0])
	}

	cmdStdOut, err := cmdRes.Output()
	if err != nil {
		fmt.Printf("%s: command not found\n", line[0])
	} else {
		fmt.Print(string(cmdStdOut))
	}
}
