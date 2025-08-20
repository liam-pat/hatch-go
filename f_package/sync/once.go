package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	oncefun := func() { fmt.Println("exec this func") }

	for i := 0; i < 10; i++ {
		go func() { once.Do(oncefun) }()
	}

	time.Sleep(1 * time.Second)
}
