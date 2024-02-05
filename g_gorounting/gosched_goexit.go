package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	{
		for i := 1; i <= 10; i++ {
			go func(i int) {
				if i == 5 {
					runtime.Gosched() // don't run first forever
				}
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
	}
	{
		for i := 1; i <= 5; i++ {
			defer fmt.Println("defer ", i)
			go func(i int) {
				if i == 3 {
					runtime.Goexit() // exit current goroutine
				}
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
	}
}
