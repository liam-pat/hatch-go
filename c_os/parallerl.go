package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
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
				fmt.Println("get the cpu no : ", <-sem)
			}
		}()
		fmt.Println(strings.Repeat("##", 30))
	}
}
