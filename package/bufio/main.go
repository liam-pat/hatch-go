package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

const tmpDir = "../../tmp"

func main() {
	{
		// write
		file, _ := os.OpenFile(tmpDir+"/bufioz.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
		defer file.Close()
		writer := bufio.NewWriter(file)

		length, _ := writer.WriteString("hello world! (write by bufio)\n")
		writer.Flush()
		fmt.Println("write length : ", length)
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		//read
		file, _ := os.OpenFile(tmpDir+"/bufioz.txt", os.O_RDONLY, os.ModePerm)
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadBytes('\n')
			line = bytes.TrimRight(line, "\r\n")

			fmt.Println(string(line))
			if err == io.EOF {
				break
			}
		}
		fmt.Println(strings.Repeat("*#", 10))

		_ = os.Remove(tmpDir + "/bufioz.txt")
	}

}
