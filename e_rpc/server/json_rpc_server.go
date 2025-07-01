package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type LittleNumber int

func (t *LittleNumber) Factorial(args *int, reply *int) error {
	*reply = 1
	for i := 2; i <= *args; i++ {
		*reply *= i
	}
	return nil
}

func main() {
	rpc.Register(new(LittleNumber))

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1235")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
