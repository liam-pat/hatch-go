package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	con, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	defer con.Close()

	// wait handled data from sever
	go func() {
		str := make([]byte, 1024)
		for {
			// get the input string
			n, err0 := os.Stdin.Read(str)
			if err0 != nil {
				fmt.Println("err =", err)
				return
			}

			// send data to serve
			_, err1 := con.Write(str[:n])
			if err1 != nil {
				fmt.Println("err =", err)
				return
			}
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err1 := con.Read(buf)
		if err1 != nil {
			fmt.Println("err = ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}

}
