package main

import (
	"fmt"
)

func fibonacci(numberCh chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case numberCh <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag:", flag)
			return
		}
	}
}
func main() {
	numberCh := make(chan int)
	quit := make(chan bool)

	//新建协程，消费者，读取 channel 的内容
	go func() {
		for i := 0; i < 20; i++ {
			number := <-numberCh
			fmt.Println(number)
		}
		quit <- true
	}()

	fibonacci(numberCh, quit)
}
