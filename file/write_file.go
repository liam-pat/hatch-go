package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 1. write directly
	file, _ := os.OpenFile("./file/write_file.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	n, _ := file.Write([]byte("hello world"))
	fmt.Printf("write number : %d\n", n)

	// 2. buffer reader
	writer := bufio.NewWriter(file)
	n1, _ := writer.WriteString("hello world!")
	fmt.Printf("write number : %d\n", n1)
	writer.Flush()

	// 3. ioutil
	s := "hello world! 123"
	ioutil.WriteFile("./file/write_file.txt", []byte(s), os.ModePerm)

	// delete file
	os.Remove("./file/write_file.txt")
}
