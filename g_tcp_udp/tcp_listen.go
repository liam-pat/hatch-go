package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// nc 127.0.0.1 3000
func main() {
	if listener, err := net.Listen("tcp", "127.0.0.1:3000"); err == nil {
		defer listener.Close()

		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Println("listener accept err:", err)
				return
			} else {
				go handler(conn)
			}
		}

	} else {
		fmt.Println("net listen err:", err)

		return
	}
}

func handler(conn net.Conn) {
	if conn == nil {
		panic("conn is nil")
	}
	defer conn.Close()

	buf := make([]byte, 4096)

	for {
		conn.Write([]byte(strings.Repeat("##", 20) + "  server -> plz input message: "))

		n, err := conn.Read(buf)

		if err == io.EOF {
			fmt.Println("conn.read to EOF")
			break
		}
		if err != nil {
			fmt.Println("conn.read get the err:", err)
			break
		}

		fmt.Println("Server receive msg:", string(buf[:n]))

		conn.Write([]byte(strings.Repeat("##", 20) + "  server -> okay, got it\n\n"))
	}
}
