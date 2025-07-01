package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, con net.Conn) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err1 := file.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("file read finish")
			} else {
				fmt.Println("read file err = ", err1)
			}
			break
		}
		_, err2 := con.Write(buf[:n])
		if err2 != nil {
			fmt.Println("send file content to server error")
		}
	}

}

func main() {

	fmt.Println("input file path , please")
	var path string
	fmt.Scan(&path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("get file detail err = ", err)
		return
	}

	con, err1 := net.Dial("tcp", "127.0.0.1:8001")
	if err1 != nil {
		fmt.Println("dial err = ", err)
		return
	}

	defer con.Close()

	_, err2 := con.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("send file name err = ", err2)
		return
	}

	buf := make([]byte, 1024)

	n, err4 := con.Read(buf)
	if err4 != nil {
		fmt.Println("send file name err = ", err4)
		return
	}

	if "ok" == string(buf[:n]) {
		sendFile(path, con)
	}

}
