package main

import (
	"fmt"
	"time"
)

func main() {
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
		time.Sleep(2 * time.Second)
	}
	<-quit
	fmt.Println("the process end")
}
