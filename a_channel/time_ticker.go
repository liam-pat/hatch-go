package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		// ticker, retry the failed task
		ticker := time.NewTicker(1 * time.Millisecond * 200)
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
		// set the timeout
		ticker := time.NewTicker(time.Millisecond * 500)
		stopper := time.NewTimer(time.Second * 2)

		var i int
		for {
			select {
			case <-stopper.C:
				fmt.Println(">>>>>>no time la, need to  stop")
				return
			case <-ticker.C:
				i++
				fmt.Println("tick times : ", i)
			}
		}
	}
}
