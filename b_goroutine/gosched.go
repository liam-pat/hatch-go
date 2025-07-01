package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	{
		for i := 1; i <= 10; i++ {
			go func(i int) {
				if i == 5 {
					runtime.Gosched() // don't run first forever
				}
				fmt.Printf("==%d", i)
			}(i)
		}
		time.Sleep(time.Second)
		fmt.Println()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Printf("-- %d", i)
			}
		}()
		for i := 0; i < 2; i++ {
			runtime.Gosched() //
			fmt.Println()
			fmt.Println("main gosched")
		}
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		for i := 1; i <= 5; i++ {
			go func(i int) {
				if i == 3 {
					fmt.Println("goroutine = ", i, "exit")
					runtime.Goexit() // exit current goroutine
				}
				fmt.Println("goroutine = ", i)
			}(i)
		}
		runtime.Gosched()
		fmt.Println(strings.Repeat("##", 20))
	}
}
