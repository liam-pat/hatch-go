package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("child")
		}
	}()

	for i := 0; i < 2; i++ {
		//先执行完子协程再执行主协程,交出控制权
		runtime.Gosched()
		fmt.Println("hello")
	}
}
