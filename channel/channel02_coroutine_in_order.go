package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)

	defer fmt.Println("main coroutine end")

	go func() {
		defer fmt.Println("child coroutine end")

		for i := 0; i < 10; i++ {
			fmt.Println("child coroutine i =", i)
			time.Sleep(time.Second)
		}
		channel1 <- "the child coroutine is end, start to run the main coroutine"
	}()

	//waiting for child coroutine
	str := <-channel1

	fmt.Println("channel1 string is :", str)
}
