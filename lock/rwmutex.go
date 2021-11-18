package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var rwm sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("Try Lock reading i:", i)
			rwm.RLock()

			time.Sleep(time.Second * 2)

			fmt.Println("Try Unlock reading i: ", i)
			rwm.RUnlock()
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try Lock writing ")
	rwm.Lock()
	fmt.Println("Ready Locked writing ")
}