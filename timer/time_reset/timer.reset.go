package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(1 * time.Second)

	<-timer.C

	fmt.Println("time is over")
}
