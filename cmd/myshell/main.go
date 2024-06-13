package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// write to standard output
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd = strings.TrimSpace(cmd)
	fmt.Printf("%s: command not found\n", cmd)
}
