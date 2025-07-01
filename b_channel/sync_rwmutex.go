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
				log.Println("got the read lock,num is", i)
				time.Sleep(time.Second * 1) // simulate the read file action
			}(i)
		}

		time.Sleep(time.Millisecond * 100) // avoid the goroutine is not ready
		rwm.Lock()
		defer rwm.Unlock()

		log.Println("main process got the write lock")
	}
}
