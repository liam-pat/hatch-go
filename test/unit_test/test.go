package main

import (
	"fmt"
	"time"
)

func main() {
	{
		for i := 1; i <= 5; i++ {
			go func() {
				fmt.Println(i)
			}()
		}
		time.Sleep(time.Second)
	}
	fmt.Println("###")

	{
		for i := 1; i <= 5; i++ {
			go func(i int) {
				fmt.Println(i)
			}(i)
		}
		time.Sleep(time.Second)
	}
}
