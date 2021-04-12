package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 2 cpu core to run
	n := runtime.GOMAXPROCS(2)
	fmt.Println("n=", n)

	for {
		go fmt.Print(1) // another coroutine
		fmt.Print(0)    // main coroutine
	}
}
