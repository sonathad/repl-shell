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
			break
		}
		fmt.Printf("%s: command not found\n", cmd)
	}

}
