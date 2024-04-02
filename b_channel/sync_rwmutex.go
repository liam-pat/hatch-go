package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	{
		var rwm sync.RWMutex

		for i := 0; i < 3; i++ {
			go func(i int) {
				rwm.RLock()
				defer rwm.RUnlock()
				time.Sleep(time.Second * 1) // simulate the read file action
				log.Println("read lock :", i)
			}(i)
		}

		time.Sleep(time.Millisecond * 100) // avoid the goroutine is not ready
		rwm.Lock()
		log.Println("get the write lock")
	}
}
