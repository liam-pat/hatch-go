package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	wait : will unlock , when the signal arrived will lock again
	signal : send signal to which is waiting
    Broadcast: broadcast to all channel
*/
func main() {

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
}
