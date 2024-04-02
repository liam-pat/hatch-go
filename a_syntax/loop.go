package main

import (
	"fmt"
	"time"
)

func main() {
	//foreach sentence
	str := "BiYongYao"
	for key, value := range str {
		fmt.Printf("key: %d ; value: %c \n", key, value)
	}

	i := 1
	for {
		time.Sleep(time.Second * 2)

		if i == 8 {
			break
		}
		fmt.Println("i=", i)
		i++
	}
}
