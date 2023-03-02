package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	{
		var a atomicInt
		a.increment()
		go func() {
			a.increment()
		}()

		fmt.Println("main start sleep 2s")
		time.Sleep(time.Second * 1)
		fmt.Println(a.get())

		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		// notice : diff lock
		var rwm sync.RWMutex
		for i := 0; i < 3; i++ {
			go func(i int) {
				rwm.RLock()
				fmt.Printf("---[time = %s] i = %d get the lock ,start sleep 2s\n", time.Now().Format("2023-01-01 15:04:05"), i)
				time.Sleep(time.Second * 2)
				rwm.RUnlock()
				fmt.Printf("---[time = %s] i = %d unlock\n", time.Now().Format("2023-01-01 15:04:05"), i)
			}(i)
		}

		time.Sleep(time.Millisecond * 100)

		fmt.Println("main process try to get Lock")
		rwm.Lock()
		fmt.Println("main process get lock successfully")
	}
}
