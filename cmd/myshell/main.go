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
	reader := bufio.NewReader(os.Stdin)

	for {
		checkCmd(reader)
	}
}

func checkCmd(reader *bufio.Reader) {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error receiving input")
		os.Exit(1)
	}

	line := strings.Split(strings.TrimSpace(input), " ")
	suf := strings.Join(line[1:], " ")

	switch pre := line[0]; pre {
	case "exit":
		os.Exit(0)
	case "":
		return
	case "echo":
		fmt.Println(strings.TrimSpace(suf))
	case "type":
		existingBuiltins := map[string]bool{
			"echo": true,
			"exit": true,
			"type": true,
			"cd":   true,
		}
		if existingBuiltins[suf] {
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
