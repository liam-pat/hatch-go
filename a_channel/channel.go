package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		// sample : common channel usage
		var channel = make(chan int)
		go func() {
			fmt.Println("exec go routine person1 start")
			channel <- 666
			fmt.Println("exec go routine person1 end")
		}()
		go func() {
			fmt.Println("exec go routine person2 start")
			<-channel
			fmt.Println("exec go routine person2 end")
		}()

		time.Sleep(time.Second * 3)
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		// sample : get data from channel
		channel := make(chan string)
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Println("child coroutine i =", i)
				time.Sleep(time.Second)
			}
			channel <- "the child coroutine is end, start to run the main coroutine"
		}()
		fmt.Println(<-channel)
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		// sample : no cache channel
		noCacheChannel := make(chan int, 0)
		go func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("---goroutine sends to no-cache-channel : i = %d \n", i)
				noCacheChannel <- 666
			}
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("main process sleeps 2s and then get no-cache-channel data")
		for i := 0; i < 3; i++ {
			fmt.Printf("main process gets no-cache-channel data : %d \n", <-noCacheChannel)
		}
		fmt.Println(strings.Repeat("*#", 30))
	}
	{
		// sample : close the channel
		noCacheChannel := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				noCacheChannel <- i
			}
			// close the channel, cannot send data anymore but still can get data
			close(noCacheChannel)
		}()
		for {
			if num, ok := <-noCacheChannel; ok == true {
				fmt.Println("get the channel data number : ", num)
			} else {
				fmt.Println("exit,because the channel is closed")
				break
			}
		}
		fmt.Println(strings.Repeat("*#", 30))
	}
	{
		// sample : producer and customer
		ch := make(chan int)
		go func(writer chan<- int) {
			for i := 0; i < 10; i++ {
				fmt.Println("write the number : ", i*i)
				writer <- i * i
			}
			close(writer)

		}(ch)
		go func(reader <-chan int) {
			for num := range reader {
				fmt.Println("reader the number", num)
			}
		}(ch)

		time.Sleep(time.Second * 2)
	}
}
