package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 3)
	ch := time.After(time.Second * 5)

	go func() {
		<-timer.C
		fmt.Println("3s later")
	}()

	go func() {
		<-ch
		fmt.Println("5s later")
	}()

	time.Sleep(time.Second * 10)
	flag := timer.Stop()
	fmt.Println(flag)

}
