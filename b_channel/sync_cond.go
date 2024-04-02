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
			cond.Signal()
		}()

		cond.L.Lock()
		defer cond.L.Unlock()

		for condition == false {
			cond.Wait()
			fmt.Println("get the goroutine signal")
		}

		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// producer vs consumer, depend on the cond send signal
		const length = 5

		rand.New(rand.NewSource(time.Now().UnixNano()))
		ch := make(chan int, length)

		var cond *sync.Cond = sync.NewCond(&sync.Mutex{})

		for i := 0; i < 5; i++ {
			go func(ch chan<- int, num int) {
				for {
					cond.L.Lock()
					for len(ch) == length { // cannot if
						cond.Wait()
					}
					data := rand.Intn(1000)
					ch <- data
					fmt.Printf(">>>> goroutine %d push -> %d\n", num, data)
					cond.L.Unlock()
					cond.Signal()
					time.Sleep(time.Second)
				}
			}(ch, i)
		}

		for i := 0; i < 5; i++ {
			go func(ch <-chan int, num int) {
				for {
					cond.L.Lock()
					for len(ch) == 0 { // cannot if
						cond.Wait()
					}
					fmt.Printf("goroutine %d receive -> %d\n", num, <-ch)
					cond.L.Unlock()
					cond.Signal()
					time.Sleep(time.Second)
				}
			}(ch, i)
		}

		select {}
	}
}
