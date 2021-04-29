package main

import (
	"fmt"
	"time"
)

func createWorker1(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("work %d recevicer %c \n", id, <-c)
		}
	}()
	return c
}

func channelDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker1(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}