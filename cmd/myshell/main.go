package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// write to standard output
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
