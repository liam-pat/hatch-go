package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./README.md"
	if content, err := ioutil.ReadFile(fileName); err == nil {
		fmt.Println(string(content))
	} else {
		fmt.Println("cannot print file content:", err)
	}
}
