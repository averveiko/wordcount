package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	text, ok := readFromArgument()
	if !ok {
		fmt.Println("usage: ./wordcount TEXT")
		return
	}

	fmt.Println(len(strings.Fields(text)))
}

func readFromArgument() (string, bool) {
	if len(os.Args) < 2 {
		return "", false
	}

	return os.Args[1], true
}
