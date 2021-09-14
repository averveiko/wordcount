package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: ./wordcount TEXT")
		return
	}

	fmt.Println(getWordCount(os.Args[1]))
}

func getWordCount(text string) int {
	return len(strings.Fields(text))
}
