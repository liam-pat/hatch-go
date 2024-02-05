package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeFile(fileName string) {
	file, _ := os.Create(fileName)
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 30; i++ {
		fmt.Fprintf(writer, "strart to write %d \n", i)
	}
}

func WriteFile(path string) {
	file, _ := os.Create(path)
	defer file.Close()

	for i := 1; i <= 10; i++ {
		str := fmt.Sprintf("line : %d \n", i)
		file.WriteString(str)
	}
}

func ReadFile(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	buffer := make([]byte, 1024*2) // create 2KB to store

	file.Read(buffer)
	fmt.Println("buf =", string(buffer[:]))
}

func ReadFileByLine(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	r := bufio.NewReader(file)

	for lineCount := 1; ; lineCount++ {
		lineString, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("line %d : %s\n", lineCount, string(lineString))
	}
}

func main() {
	{
		writeFile("./file/start_to_write.txt")

		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		path := "./file/write_to_read.txt"
		WriteFile(path)

		ReadFile(path)
		ReadFileByLine(path)
	}
}
