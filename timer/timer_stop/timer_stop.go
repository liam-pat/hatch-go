package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("child process is running")
	}()

	//因为timer已经停止了，所以子协程一直阻塞
	timer.Stop()

	for {

	}
}
