package main

import (
	"fmt"
	"runtime"
)

func testDefer() {
	defer fmt.Println("C")
	//return
	// 终止函数，所在的子协程还会继续执行，B 还是会输出

	runtime.Goexit() // 终止函数所在的协程，B 不会再输出了
	fmt.Println("D")
}

func main() {

	go func() {
		fmt.Println("A")
		testDefer()
		fmt.Println("B")
	}()

	select {}
}
