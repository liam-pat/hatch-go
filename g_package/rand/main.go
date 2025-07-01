package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(66)
	for i := 0; i < 5; i++ {
		fmt.Println("random number:", rand.Int())
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println("random number: ", rand.Intn(100))
	}
}
