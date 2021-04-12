package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func person1() {
	printer("hello")
	channel <- 666
}
func person2() {
	<-channel
	printer("world")
}

func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func main() {
	go person1()

	go person2()

	time.Sleep(time.Second * 20)
}
