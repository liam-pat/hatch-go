package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	parameters := os.Args

	if len(parameters) != 3 {
		fmt.Println("parameter error , double check please")

		return
	}

	srcFileName := parameters[1]
	dstFileName := parameters[2]

	if srcFileName == dstFileName {
		fmt.Println("source file name equal dst file name ")
		return
	}

	srcFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("open file got error", err1)
		return
	}

	dstFile, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("create file got error", err1)
		return
	}

	defer srcFile.Close()
	defer dstFile.Close()

	buf := make([]byte, 4*1024)

	for {
		readNum, err := srcFile.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read got error", err)
		}

		_, err3 := dstFile.Write(buf[:readNum])
		if err3 != nil {
			fmt.Println("write got error", err)
		}
	}
}
