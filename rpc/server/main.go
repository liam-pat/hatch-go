package main

import (
	"biyongyao.com/go-reborn/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcDemo.DemoService{})

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err : %v ", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
