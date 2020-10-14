package main

import (
	"fmt"
	"time"
)

func main() {

	noCacheChannel := make(chan int, 0)

	fmt.Printf("len: %d ; cap: %d \n", len(noCacheChannel), cap(noCacheChannel))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("############")
			fmt.Printf("child process: i = %d \n", i)
			noCacheChannel <- 1
			fmt.Printf("child process: len = %d ,cap = %d \n", len(noCacheChannel), cap(noCacheChannel))
		}
	}()

	//因为这里等待了2秒，子协程肯定先执行，下面的代码不会阻塞
	// 没有缓存的 channel 没取出 ,以为缓存区不存储数据，要等待下面代码取出数据，子协程的 for 循环才会继续下去
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println("$$$$$$$$$$$$")
		fmt.Printf("main process: i = %d \n", i)
		temp := <-noCacheChannel
		fmt.Printf("noCacheChannel : %d \n", temp)
	}
}
