package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx xxxx")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName {
		fmt.Println("src handle_file is same as dst handle_file")
		return
	}

	sourceFile, err1 := os.Open(srcFileName)

	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}

	newFile, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2", err2)
		return
	}
	defer sourceFile.Close()
	defer newFile.Close()

	buffer := make([]byte, 4*1024)
	for {
		n, err := sourceFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}
		writeNumber, err3 := newFile.Write(buffer[:n])

		if err3 != nil {
			fmt.Printf("have something error in %d", writeNumber)
			continue
		}
	}
}
