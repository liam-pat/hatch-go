package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	{
		fmt.Println("cpu num: ", runtime.NumCPU())
		fmt.Println("thread num: ", runtime.GOMAXPROCS(runtime.NumCPU()))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		// init the goroutine needs some time, i may already updated to 10 when the goroutine is created
		for i := 0; i < 10; i++ {
			go func() {
				fmt.Printf("==%d", i)
			}()
		}
		time.Sleep(time.Millisecond)
		fmt.Println()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		for i := 0; i < 10; i++ {
			go func(number int) {
				fmt.Printf("==%d", number)
			}(i)
		}
		time.Sleep(time.Millisecond)
		fmt.Println()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		const num = 4
		runtime.GOMAXPROCS(num)

		func() {
			sem := make(chan int, num)

			for i := 0; i < num; i++ {
				go func(sem chan int, num int) {
					sem <- num
				}(sem, i)
			}

			for i := 0; i < num; i++ {
				fmt.Println("goroutine id: ", <-sem)
			}
		}()
		fmt.Println(strings.Repeat("##", 30))
	}
}
