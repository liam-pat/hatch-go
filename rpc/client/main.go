package main

import (
	"biyongyao.com/go-reborn/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	con, err := net.Dial("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(con)

	var result float64

	err = client.Call("DemoService.Div", rpcDemo.Args{A: 30, B: 3}, &result)

	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcDemo.Args{A: 30, B: 30}, &result)

	fmt.Println(result, err)
}
