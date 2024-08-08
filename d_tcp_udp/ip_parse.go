package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("error input")
	}

	ipStr := os.Args[1]
	addr := net.ParseIP(ipStr)

	if addr == nil {
		fmt.Println("err address")
	} else {
		fmt.Printf("parse to ip: %v \n", addr)
	}
	os.Exit(0)
}
