package main

import (
	"log"
	"net/rpc"
)

type Result1 struct{ Num, Ans int }

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	var result Result1
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call remote process Cal.Square.", err)
	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
