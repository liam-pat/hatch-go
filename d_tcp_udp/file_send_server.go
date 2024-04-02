package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func receiveFile(fileName string, con net.Conn) {

	f, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Cannot create the file", err)
		return
	}

	buf := make([]byte, 1024*4)
	for {
		n, err1 := con.Read(buf)
		if err1 != nil {
			if err == io.EOF {
				fmt.Println("already get all file content", err1)
			} else {
				fmt.Println("read file content from remote fail", err1)
			}
			return
		}
		if n == 0 {
			fmt.Println("already get all file content and write to file finish , ", err1)
			return
		}
		_, err2 := f.Write(buf[:n])
		if err2 != nil {
			fmt.Println("write remote data to file fail", err)
			return
		}
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("Cannot not init listener", err)
		return
	}
	defer listener.Close()

	con, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("Cannot accept the connection", err)
		return
	}
	defer con.Close()

	buf := make([]byte, 1024)
	n, err2 := con.Read(buf)

	if err2 != nil {
		fmt.Println("Read the data error", err)
		return
	}

	filename := string(buf[:n])
	_, err3 := con.Write([]byte("ok"))
	if err3 != nil {
		fmt.Println("send the data error", err)
		return
	}

	receiveFile(filename, con)

}
