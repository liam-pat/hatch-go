package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const BUFFER = 5

var cond = sync.NewCond(&sync.Mutex{})

// producer
func producer(ch chan<- int) {
	for {
		cond.L.Lock()
		fmt.Printf("[%s] producer get the lock\n", time.Now().Format("2006-01-02 15:04:05"))

		for len(ch) == BUFFER { // if the buffer full , wait consume and notify
			cond.Wait()
		}

		ch <- rand.Intn(1000)
		time.Sleep(time.Second * 1)

		cond.L.Unlock()
		cond.Signal()
	}
}

// consumer
func consumer(ch <-chan int) {
	for {
		cond.L.Lock()
		fmt.Printf("[%s] consumer get the lock\n", time.Now().Format("2006-01-02 15:04:05"))

		for len(ch) == 0 {
			cond.Wait() // will unload , waiting the signal
		}
		fmt.Println("Receive:", <-ch)

		time.Sleep(time.Second * 1)

		cond.L.Unlock()
		cond.Signal()
	}
}
func main() {
	{
		cond := sync.NewCond(&sync.Mutex{})
		condition := false

		go func() {
			time.Sleep(time.Second * 1)
			cond.L.Lock()

			condition = true
			cond.Signal() // send signal

			cond.L.Unlock()
		}()

		cond.L.Lock()
		for !condition {
			cond.Wait() //will unlock . get the signal will lock again
			fmt.Println("get the signal")
		}
		cond.L.Unlock()
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		ch := make(chan int, BUFFER)

		for i := 0; i < 5; i++ {
			fmt.Printf("start to create %d producer goruntine \n", i)
			go producer(ch)
		}

		for i := 0; i < 5; i++ {
			fmt.Printf("start to create %d consumer goruntine \n", i)
			go consumer(ch)
		}

		select {}
	}
}
