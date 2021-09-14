package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	text, ok := readFromArgument()
	if !ok {
		fmt.Println("0")
		return
	}

	fmt.Println(len(strings.Split(text, " ")))
}

// read from command line
func readFromArgument() (string, bool) {
	if len(os.Args) < 2 {
		return "", false
	}
	if os.Args[1] == "" {
		return "", false
	}
	return os.Args[1], true
}
