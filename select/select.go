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
	numberCh := make(chan int) //无缓存

	quit := make(chan bool) //无缓存

	//新建协程，消费者，读取 channel 的内容
	go func() {
		for i := 0; i < 20; i++ {
			number := <-numberCh
			fmt.Println(number)
		}
		quit <- true
	}()

	//生产者，产生数字，写入 channel
	fibonacci(numberCh,quit)
}
