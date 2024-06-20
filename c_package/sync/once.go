package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("exec this func")
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
		}()
	}

	time.Sleep(1 * time.Second)
}
