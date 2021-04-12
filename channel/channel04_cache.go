package main

import (
	"fmt"
	"time"
)

func main() {

	cacheChannel := make(chan int, 3)

	fmt.Printf("len: %d ; cap: %d \n", len(cacheChannel), cap(cacheChannel))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("")
			fmt.Printf("child process: i = %d \n", i)
			cacheChannel <- 666
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println("")
		fmt.Printf("main process: i = %d \n", i)
		temp := <-cacheChannel
		fmt.Printf("cacheChannel : %d \n", temp)
	}
}
