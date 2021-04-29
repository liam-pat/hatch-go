package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker will send data to channel 2s
	ticker := time.NewTicker(2 * time.Second)

	i := 0

	for {
		<-ticker.C
		i++

		fmt.Printf("i=%d\n", i)

		if i == 10 {
			break
		}
	}
}