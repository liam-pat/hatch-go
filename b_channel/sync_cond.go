package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	{
		cond := sync.NewCond(&sync.Mutex{})
		condition := false

		go func() {
			time.Sleep(time.Second * 1)

			cond.L.Lock()
			defer cond.L.Unlock()

			condition = true
			cond.Signal() // send the signal to the goroutine, and then release the lock
		}()

		cond.L.Lock()
		defer cond.L.Unlock()

		for condition == false {
			cond.Wait() // will release the lock and wait for the signal
			fmt.Println("get the goroutine signal")
		}

		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// producer vs consumer, depend on the cond send signal
		rand.New(rand.NewSource(time.Now().UnixNano()))

		const length = 5
		ch := make(chan int, length)

		var cond *sync.Cond = sync.NewCond(&sync.Mutex{})

		for i := 0; i < length; i++ {
			go func(ch chan<- int, num int) {
				cond.L.Lock()
				for len(ch) == length { // cannot if
					cond.Wait()
				}
				rand := rand.Intn(1000)
				ch <- rand
				fmt.Printf(">>>>AAA goroutine %d push -> %d\n", num, rand)
				cond.L.Unlock()
				cond.Signal()
				time.Sleep(time.Millisecond * 500)
			}(ch, i)
		}

		for i := 0; i < length; i++ {
			go func(ch <-chan int, num int) {
				cond.L.Lock()
				for len(ch) == 0 { // cannot if
					cond.Wait()
				}
				fmt.Printf(">>>>BBB goroutine %d receive -> %d\n", num, <-ch)
				cond.L.Unlock()
				cond.Signal()
				time.Sleep(time.Millisecond * 500)
			}(ch, i)
		}

		time.Sleep(time.Second * 3)

	}
}
