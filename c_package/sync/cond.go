package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

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
}
