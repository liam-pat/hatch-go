package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

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
