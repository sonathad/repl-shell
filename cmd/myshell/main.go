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
		switch cmd {
		case "":
			return
		case "exit 0":
			os.Exit(0)
		}
		fmt.Printf("%s: command not found\n", cmd)
	}
}
