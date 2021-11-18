package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//set the const seed
	fmt.Println("---------const seed---------")
	rand.Seed(66)
	for i := 0; i < 5; i++ {
		fmt.Println("random number:", rand.Int())
	}
	// timestamp set to seed
	fmt.Println("---------different seed---------")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println("random number: ", rand.Intn(100))
	}
}

