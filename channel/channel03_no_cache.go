package main

import (
	"fmt"
	"time"
)

func main() {

	// no  define the cap 0 ,default 0
	noCacheChannel := make(chan int, 0)
	fmt.Printf("len: %d ; cap: %d \n", len(noCacheChannel), cap(noCacheChannel))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("")
			fmt.Printf("child process: i = %d \n", i)
			noCacheChannel <- 666
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println()
		fmt.Printf("main process: i = %d \n", i)

		temp := <-noCacheChannel
		fmt.Printf("noCacheChannel : %d \n", temp)
	}
}
