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
			for i := 0; i < 5; i++ {
				fmt.Println("child -> ", i)
			}
		}()
		for i := 0; i < 2; i++ {
			runtime.Gosched() //
			fmt.Println("hello -> ", i)
		}
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		for i := 1; i <= 10; i++ {
			go func(i int) {
				if i == 5 {
					runtime.Gosched() // no 5 will print at the end
				}
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
		fmt.Println(strings.Repeat("##", 20))
	}
}
