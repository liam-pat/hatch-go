package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleCon(con net.Conn) {
	defer con.Close()

	addr := con.RemoteAddr().String()
	fmt.Printf("addr %s connect successful", addr)

	buf := make([]byte, 2048)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("err =", err)
			return
		}
		fmt.Printf("[%s] receive original data len %d : %s \n", addr, len(string(buf[:n])), string(buf[:n]))

		if "exit" == string(buf[:n-1]) {
			fmt.Println("client input exit to break : ", addr)
			return
		}

		_, err2 := con.Write([]byte(strings.ToUpper(string(buf[:n]))))

		if err2 != nil {
			fmt.Println("err =", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer listener.Close()

	for {
		con, err1 := listener.Accept()

		if err1 != nil {
			fmt.Println("err =", err)
			return
		}

		go HandleCon(con)
	}
}
