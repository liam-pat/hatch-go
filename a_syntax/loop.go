package main

import (
	"fmt"
	"time"
)

func main() {
	str := "foreach one"
	i := 1

	number := 10000000
	better := 10_000_000

	fmt.Println(number, "-", better)

	for key, value := range str {
		fmt.Printf("key: %d ; value: %c \n", key, value)
	}
	for {
		time.Sleep(time.Second * 1)
		if i == 8 {
			break
		}
		fmt.Println("i=", i)
		i++
	}
}
