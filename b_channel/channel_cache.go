package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		messages := make(chan string, 2)
		messages <- "buffered1"
		messages <- "buffered2"
		fmt.Println("string1 : ", <-messages)
		fmt.Println("string2 : ", <-messages)

		fmt.Println(strings.Repeat("##", 30))
	}
	{
		messages := make(chan string, 2)
		go func(channel chan<- string, msg string) {
			channel <- msg
		}(messages, "message")

		fmt.Printf("output message: %s\n", <-messages)

		fmt.Println(strings.Repeat("##", 30))
	}
	{
		writer := make(chan string, 1)
		go func(writer chan<- string) {
			time.Sleep(time.Second * 1)
			writer <- "write message"
		}(writer)
		fmt.Println("wait for 1s to write message")
		<-writer
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		jobs := make(chan int, 5)
		done := make(chan bool)
		go func() {
			for {
				if j, ok := <-jobs; ok {
					fmt.Println("jobs channle opened and receive: ", j)

				} else {
					fmt.Println("jobs channle was closed")
					done <- true
					return
				}
			}
		}()
		for j := 1; j <= 3; j++ {
			time.Sleep(time.Second * 1)
			jobs <- j
		}
		close(jobs)
		<-done

		fmt.Println(strings.Repeat("##", 30))
	}
	{
		ch := make(chan string, 3)

		for i := 0; i < 20; i++ {

			go func(ch chan string) {
				ch <- "push"

				fmt.Printf("i = %d goroutine in\n", i)
				time.Sleep(time.Second * 1)

				<-ch
				fmt.Printf("i = %d goroutine out\n", i)
			}(ch)

		}
		time.Sleep(time.Second * 10)
	}
}
