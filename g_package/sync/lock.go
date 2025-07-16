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
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func main() {
	{
		var a atomicInt
		a.increment()
		go func() { a.increment() }()

		time.Sleep(time.Second * 1)
		fmt.Printf("%+ v\n", a)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		var rwm sync.RWMutex
		defer rwm.Unlock()
		for i := 0; i < 3; i++ {
			go func(i int) {
				rwm.RLock()
				time.Sleep(time.Second * 2)
				rwm.RUnlock()
			}(i)
		}

		time.Sleep(time.Millisecond * 100)
		rwm.Lock()
		fmt.Println("main process get lock successfully")
	}
}
