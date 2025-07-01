package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1235")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply int
	err = client.Call("LittleNumber.Factorial", 10, &reply)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Printf("10! = %d\n", reply)
}
