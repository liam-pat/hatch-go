package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		var channel = make(chan int)
		// can't control which one finishes first.
		go func() {
			fmt.Println("goroutine 1 start")
			channel <- 666
			fmt.Println("goroutine 1 end")
		}()
		go func() {
			fmt.Println("goroutine 2 start")
			<-channel
			fmt.Println("goroutine 2 end")
		}()

		time.Sleep(time.Second)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		channel := make(chan string)
		go func() {
			for i := 0; i < 5; i++ {
				time.Sleep(time.Millisecond)
			}
			channel <- "goroutine end"
		}()
		fmt.Println("main receive the goroutine signal .. (", <-channel)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		noCacheChannel := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				noCacheChannel <- i
			}
			close(noCacheChannel)
		}()
		for {
			if num, ok := <-noCacheChannel; ok == true {
				fmt.Println("receive data : ", num)
			} else {
				fmt.Println("channel was closed")
				break
			}
		}
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// if the channel is not closed, the goroutine will be blocked to send data
		// but can still get the data from the channel
		ch := make(chan int)
		done := make(chan bool)
		go func(writer chan<- int) {
			for i := 0; i < 5; i++ {
				writer <- i * i
			}
			close(writer)
		}(ch)

		go func(reader <-chan int) {
			for num := range reader {
				fmt.Println("goroutine 2 receive data:", num)
			}
			done <- true
		}(ch)

		<-done
		fmt.Println(strings.Repeat("##", 30))
	}
}
