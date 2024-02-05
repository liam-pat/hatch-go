package main

import (
	"fmt"
	"time"
)

func main() {
	{
		// the main process is ended, and the i is set to 11, goroutine
		for i := 1; i <= 5; i++ {
			go func() {
				fmt.Println(i)
			}()
		}
		time.Sleep(time.Second)
	}

	fmt.Println("###")

	{
		// diff goroutines exec time is diff
		for i := 1; i <= 5; i++ {
			go func(i int) {
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
	}
}
