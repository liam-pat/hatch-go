package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		messages := make(chan string, 2)
		messages <- "buffered 01"
		messages <- "buffered 01"
		fmt.Println("string 1 : ", <-messages)
		fmt.Println("string 2 : ", <-messages)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		messages := make(chan string, 2)
		go func(channel chan<- string, msg string) {
			channel <- msg
		}(messages, "message")
		fmt.Println(<-messages)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		// block
		writer := make(chan string, 1)
		go func(writer chan<- string) {
			time.Sleep(time.Millisecond * 500)
			writer <- "write message"
		}(writer)
		<-writer
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		jobs := make(chan int, 5)
		done := make(chan bool)
		go func() {
			for {
				if j, ok := <-jobs; ok {
					fmt.Println("goroutine receive data : ", j)

				} else {
					fmt.Println("channel was closed")
					done <- true
					return
				}
			}
		}()
		for j := 1; j <= 3; j++ {
			jobs <- j
		}
		close(jobs)
		<-done
	}
}
