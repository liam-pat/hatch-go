package main

import (
	"fmt"
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
		locker := new(sync.Mutex)
		cond := sync.NewCond(locker)

		done := false
		cond.L.Lock()

		go func() {
			time.Sleep(2e9)
			done = true
			cond.Signal()
		}()

		if !done {
			cond.Wait()
			fmt.Println("wait the go routine done")
		}

		fmt.Println("go routine is done : ", done)
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		var locker = new(sync.Mutex)
		var cond = sync.NewCond(locker)

		for i := 0; i < 5; i++ {
			go func(x int) {
				cond.L.Lock()
				cond.Wait()
				fmt.Printf("no %d gorountine resumed \n", x)
				cond.L.Unlock()
			}(i)
		}

		time.Sleep(2 * time.Second)
		fmt.Println("main process broadcast all")
		cond.Broadcast()
		time.Sleep(2 * time.Second)
	}
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
			cond.Wait()
			fmt.Println("wait the go routine signal")
		}
		cond.L.Unlock()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		ch := make(chan int, BUFFER)
		for i := 0; i < 5; i++ {
			go producer(ch)
		}

		for i := 0; i < 5; i++ {
			go consumer(ch)
		}

		select {}
	}
}
