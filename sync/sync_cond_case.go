package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const BUFFER = 5

var cond = sync.NewCond(&sync.Mutex{})

// producer
func producer(ch chan<- int) {
	for {
		cond.L.Lock()

		for len(ch) == BUFFER {  // if the buffer full , wait consume and notify
			cond.Wait()
		}

		ch <- rand.Intn(1000)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second * 2)
	}
}

// consumer
func consumer(ch <-chan int) {
	for {
		cond.L.Lock()
		for len(ch) == 0 {
			cond.Wait()  // wait producer
		}
		fmt.Println("Receive:", <-ch)

		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Second * 1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, BUFFER)

	for i := 0; i < 10; i++ {
		go producer(ch)
	}

	for i := 0; i < 10; i++ {
		go consumer(ch)
	}

	select {}
}