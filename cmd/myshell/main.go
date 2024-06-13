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
		if cmd == "" {
			return
		} else if cmd == "exit 0" {
			os.Exit(0)
		} else if strings.HasPrefix(cmd, "echo ") {
			res, _ := strings.CutPrefix(cmd, "echo")
			fmt.Println(strings.TrimSpace(res))
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
