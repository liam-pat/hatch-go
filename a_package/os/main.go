package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const tmpDir = "../../tmp"

func main() {
	{
		// write
		file, _ := os.Create(tmpDir + "/test.txt")
		defer file.Close()

		f, _ := os.OpenFile(tmpDir+"/test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
		defer f.Close()

		byteNum, _ := f.Write([]byte("hello (by []byte)\n"))
		fmt.Println("write byte[] number : ", byteNum)

		fileInfo, _ := os.Stat(tmpDir + "/test.txt")
		fmt.Printf("name : %s \tmode : %s, \tsize : %d B \n", fileInfo.Name(), fileInfo.Mode(), fileInfo.Size())
	}
	{
		absPath, _ := filepath.Abs(tmpDir + "/test.txt")
		fmt.Println("abs path : ", absPath)

		//read
		f, _ := os.Open(absPath)
		defer f.Close()
		readByte := make([]byte, 128)
		for {
			n, _ := f.Read(readByte)
			fmt.Println(strings.Repeat("**", 5) + string(readByte[:n]))
			if n < 128 {
				fmt.Println("read end")
				break
			}
		}
		_ = os.Remove(absPath)
	}
}
