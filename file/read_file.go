package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	FILE_NAME = "./file/data.csv"
)

func main() {
	/**
	1. read directly
	*/
	f, _ := os.Open(FILE_NAME)
	defer f.Close()

	readByte := make([]byte, 128)
	for {
		n, err := f.Read(readByte)
		// read error
		if err != nil && err != io.EOF {
			fmt.Println("read err : ", err)
			break
		}
		fmt.Println(string(readByte[:n]))
		// already finish
		if n < 128 {
			fmt.Println("read end")
			break
		}
	}

	fmt.Println("###############################################################")

	/**
	2. bufio reader
	*/
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString(',') // ready one by one
		if err != nil && err != io.EOF {
			fmt.Println("read err: ", err)
			break
		}
		fmt.Println(str)
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
	}

	fmt.Println("########################################################################")

	/**
	3. ioutil reader
	*/
	ret, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		fmt.Println("read err :", err)
		return
	}
	fmt.Println(string(ret))

}
