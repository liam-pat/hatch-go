package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer file.Close()

	var buffer string

	for i := 1; i <= 10; i++ {
		buffer = fmt.Sprintf("已经写完,第%d行\n", i)

		_, err := file.WriteString(buffer)

		if err != nil {
			fmt.Println("err", err)
		}
	}
}

func ReadFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 1024*2)

	n, err1 := file.Read(buffer)

	if err1 != nil && err1 != io.EOF {
		fmt.Println("err", err)
		return
	}
	fmt.Println("buf =", string(buffer[:n]))
}

func ReadFileByLine(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for lineCount := 1; ; {
		lineString, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err", err)
		}

		fmt.Printf("line %d : %s", lineCount, string(lineString))
		lineCount++
	}
}

func main() {
	path := "./test_write_file.txt"
	WriteFile(path)
	//ReadFile(path)
	ReadFileByLine(path)
}
