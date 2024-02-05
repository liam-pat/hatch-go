package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Cal int

type Result struct{ Num, Ans int }

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	_ = rpc.Register(new(Cal)) // will detect all fun is conform
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
