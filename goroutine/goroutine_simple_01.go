package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// the result is unstable.
		// if it's a good cpu , will print some number < 10
		// if not , 10 times will print the 10
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond)
}
