package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	{
		wait := make(chan bool, 1)
		go func() {
			func() {
				defer func() {
					wait <- true
				}()
				fmt.Println("exit immediately")
				runtime.Goexit()
			}()
		}()

		<-wait
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		const NUM = 5
		wait := make(chan bool, NUM)

		for i := 1; i <= NUM; i++ {
			go func(i int) {
				defer func() { wait <- true }()
				if i == 3 {
					runtime.Goexit()
				}
				fmt.Println("goroutine number", i)
			}(i)
		}
		for i := 0; i < NUM; i++ {
			<-wait
		}
		fmt.Println(strings.Repeat("##", 20))
	}
}
