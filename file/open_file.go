package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// create new file
	file, err := os.Create("./file/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// output the abs path
	absPath, _ := filepath.Abs("./file/test.txt")
	fmt.Println("abs path: ", absPath)

	// open file
	f, err := os.OpenFile("./file/test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	// write byte[] to file
	byteNum, err := f.Write([]byte("hello"))
	if err != nil {
		fmt.Println("write byte[] err: ", err)
		return
	}
	fmt.Println("write byte[] number : ", byteNum)

	// write string to file
	stringNum, err := f.WriteString("hello")
	if err != nil {
		fmt.Println("write string err: ", err)
		return
	}
	fmt.Println("write string number = ", stringNum)

	// get file information
	fileInfo, err := os.Stat("./file/test.txt")
	if err != nil {
		fmt.Println("get file info error: ", err)
		return
	}
	fmt.Printf("name : %s ,mode : %s, size : %d B", fileInfo.Name(), fileInfo.Mode(), fileInfo.Size())

	// delete file
	err = os.Remove("./file/test.txt")
	if err != nil {
		fmt.Println("remove err:", err)
		return
	}

}
