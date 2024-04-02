package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Params struct {
	A, B int
}

type Q struct {
	Quo, Rem int
}

type Number int

func (t *Number) Multiply(args *Params, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Number) Divide(args *Params, quo *Q) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	num := new(Number)
	rpc.Register(num)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("error ", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("error ", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
