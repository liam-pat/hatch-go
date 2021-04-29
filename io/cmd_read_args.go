package main

import (
	"os"
	"strings"
)

func main() {
	var inputStr string
	if len(os.Args) > 1 {
		inputStr += strings.Join(os.Args[1:], " ")
	}
}
