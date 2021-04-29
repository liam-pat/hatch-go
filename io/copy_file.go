package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourceFileName := "./io/file/source.txt"
	dstFileName := "./io/file/copy.txt"

	_, err := CopyFile(sourceFileName, dstFileName)
	if err != nil {
		fmt.Sprintln("err : ", err)
	}

	fmt.Println("Copy done!")
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
