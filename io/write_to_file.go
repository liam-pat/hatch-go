package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// write file if the file ont exist will create , overwrite
	outputFile, outputError := os.OpenFile("./io/file/write_01.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world! golang , I love you\n"
	for i := 0; i < 10; i++ {
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			continue
		}
		fmt.Printf("Already write %d times \n", i+1)
	}
	outputWriter.Flush()
}
