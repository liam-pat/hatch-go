package main

import (
	"fmt"
	"time"
)

func main() {
	// will print the un regular number , each thread initialization speed cannot evaluate
	for i := 0; i < 10; i++ {
		go func(number int) {
			fmt.Println(number)
		}(i)
	}
	time.Sleep(time.Millisecond)
}
