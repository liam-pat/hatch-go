package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	A, B int
}

type Q struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1236")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Params{17, 8}
	var reply int
	err = client.Call("Number.Multiply", args, &reply)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Q
	err = client.Call("Number.Divide", args, &quot)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Printf("%d / %d = %d ... %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
