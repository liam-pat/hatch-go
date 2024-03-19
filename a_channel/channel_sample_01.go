package main

import (
	"fmt"
	"time"
)

func main() {
	// concurrency limitation
	ch := make(chan string, 5) // gt 5 . will be blocked
	for i := 0; i < 100; i++ {
		go func(ch chan string) {
			ch <- "push"

			fmt.Println("task push")
			time.Sleep(time.Second * 1)

			<-ch
		}(ch)
	}

	time.Sleep(time.Second * 10)
}
