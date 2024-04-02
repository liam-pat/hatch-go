package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234"); err != nil {
		log.Fatal("dialing:", err)
	} else {
		var reply int
		args := Args{17, 8}

		err = client.Call("Num.Multiply", args, &reply)
		if err != nil {
			log.Fatal("error:----", err)
		}
		fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

		var quot Quotient
		err = client.Call("Num.Divide", args, &quot)
		if err != nil {
			log.Fatal("error:---", err)
		}
		fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
	}
}
