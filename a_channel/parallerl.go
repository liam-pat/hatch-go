package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	{
		const NUMOFCPU = 4
		runtime.GOMAXPROCS(NUMOFCPU)
		func() {
			sem := make(chan int, NUMOFCPU)
			for i := 0; i < NUMOFCPU; i++ {
				go func(sem chan int, num int) {
					sem <- num
				}(sem, i)
			}
			for i := 0; i < NUMOFCPU; i++ {
				fmt.Println("Get the CPU number : ", <-sem)
			}
		}()
		fmt.Println(strings.Repeat("*#", 30))
	}
}
