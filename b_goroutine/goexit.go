package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	{
		go func() {
			fmt.Println("goroutine A")
			func() {
				defer fmt.Println("goroutine defer C")
				runtime.Goexit() // exit immediately
				fmt.Println("goroutine D")
			}()

			fmt.Println("goroutine B")
		}()

		time.Sleep(time.Millisecond)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		for i := 1; i <= 5; i++ {
			defer fmt.Println("defer ", i)

			go func(i int) {
				if i == 3 {
					runtime.Goexit()
				}
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
		fmt.Println(strings.Repeat("##", 20))
	}
}
