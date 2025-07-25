package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i <= 15; i++ {
			number := <-ch
			fmt.Printf("fibonacci(%[1]d) = %d\n", i, number)
		}
		quit <- true
	}()

	func(ch chan<- int, quit <-chan bool) {
		x, y := 0, 1
		for {
			select {
			case ch <- x:
				x, y = y, x+y
			case <-quit:
				return
			}
		}
	}(ch, quit)
}
