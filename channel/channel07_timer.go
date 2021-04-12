package main

import (
	"fmt"
	"time"
)

func main() {

	// after 10s  , it will writer time to timer.c
	timer := time.NewTimer(10 * time.Second)

	fmt.Println("current time:", time.Now())

	afterTime := <-timer.C

	fmt.Println("after time:", afterTime)
}
