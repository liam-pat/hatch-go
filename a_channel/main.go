package main

import (
	"fmt"
	"strings"
	"time"
)

func pull(channel <-chan string) string {
	return <-channel
}
func push(channel chan<- string, msg string) {
	channel <- msg
}

func do1Thing(writeMsg chan<- string) {
	fmt.Println("-----go routine start")
	time.Sleep(time.Second)
	fmt.Println("-----go routine end")
	writeMsg <- "write test msg"
}

func main() {
	{
		// simple channel buffer
		messages := make(chan string, 2)

		messages <- "buffered"
		messages <- "channel"

		fmt.Println("channel string 1 : ", <-messages)
		fmt.Println("channel string 2 : ", <-messages)
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		// common channel
		strChan := make(chan string, 2)
		push(strChan, "test to get channel msg")
		fmt.Println(pull(strChan))
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		// channel exchange
		writer := make(chan string, 1)

		go do1Thing(writer)

		fmt.Println("main process waiting")
		<-writer
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		jobs := make(chan int, 5)
		done := make(chan bool)

		go func() {
			for {
				j, got := <-jobs
				if got {
					fmt.Println("----received job : ", j)
				} else {
					fmt.Println("----cannot received maybe job has done")
					done <- true
					return
				}
			}
		}()

		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		close(jobs)
		fmt.Println("sent 3 jobs and then close")

		<-done
	}
}
