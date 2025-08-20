package main

import (
	"fmt"
	"io"
	"net"
)

// nc 127.0.0.1 3000
func main() {

	if listener, err := net.Listen("tcp", "127.0.0.1:3000"); err == nil {
		defer listener.Close()
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Println("accept err:", err)
				return
			} else {
				go handler(conn)
			}
		}

	} else {
		fmt.Println("listen err: ", err)
		return
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	buf := make([]byte, 4096)

	for {
		conn.Write([]byte("send to server: "))

		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("conn.read EOF")
			break
		}
		if err != nil {
			fmt.Println("conn.read err: ", err)
			break
		}
		fmt.Printf("%s: %s\n", remoteAddr, string(buf[:n]))
		conn.Write([]byte(fmt.Sprintf("--------\n")))
	}
}
