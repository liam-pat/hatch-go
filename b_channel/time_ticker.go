package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		// ticker, retry the failed task
		ticker := time.NewTicker(time.Millisecond * 200)
		defer ticker.Stop()

		i := 0
		for {
			<-ticker.C
			i++
			fmt.Printf("times = %d\n", i)
			if i == 3 {
				fmt.Println("retry time is maximum")
				break
			}
		}
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// ticker vs timer
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()

		timer := time.NewTimer(time.Second * 3)
		defer timer.Stop()

		var i int
		for {
			select {
			case <-timer.C:
				fmt.Println("3s la ~")
				return
			case <-ticker.C:
				i++
				fmt.Println("ticker", i)
			}
		}
	}
}
