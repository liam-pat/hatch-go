package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Param1 struct {
	A, B int
}

type Q1 struct {
	Quo, Rem int
}

func main() {

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Param1{17, 8}
	var reply int
	err = client.Call("Num1.Multiply", args, &reply)
	if err != nil {
		log.Fatal("error:----", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Q1
	err = client.Call("Num1.Divide", args, &quot)
	if err != nil {
		log.Fatal("error:----", err)
	}
	fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
