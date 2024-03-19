package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		c1 := make(chan string, 1)
		c2 := make(chan string, 1)
		go func() {
			time.Sleep(2 * time.Second)
			c1 <- "result 1"
		}()
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "result 2"
		}()
		select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
		}

		fmt.Println(strings.Repeat("*#", 30))
	}
	{
		ch := make(chan int)
		quit := make(chan bool)

		go func() {
			for {
				select {
				case num := <-ch:
					fmt.Println("num = ", num)
				case <-time.After(5 * time.Second):
					fmt.Println("overtime , 5 second after")
					quit <- true
				}
			}
		}()
		for i := 0; i < 8; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		<-quit
		fmt.Println(strings.Repeat("*#", 30))
	}
	{
		// 3 ways to wait
		way1 := func() {
			time.Sleep(1 * time.Second)
			fmt.Println("-- way1 done")
		}
		way2 := func() {
			timer := time.NewTimer(1 * time.Second)

			fmt.Println("-- way2 done ", <-timer.C)
		}
		way3 := func() {
			fmt.Println("-- way3 done", <-time.After(time.Second))
		}
		way1()
		way2()
		way3()

		fmt.Println(strings.Repeat("*#", 30))
	}
}
