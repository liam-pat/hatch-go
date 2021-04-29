package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("./io/file/test.json")
	buf := make([]byte, 1024) // you can read from http content or read from file

	inputReader := bufio.NewReader(file)

	var content string
	for {
		n, _ := inputReader.Read(buf)
		if n == 0 {
			fmt.Println("All content already read")
			break
		}
		content += string(buf[:n])
	}

	fmt.Println(content)
}
