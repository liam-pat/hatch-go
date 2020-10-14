package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//set the const seed
	rand.Seed(66)
	fmt.Println("const seed")
	for i := 0; i < 5; i++ {
		fmt.Println("random number:", rand.Int())
	}

	fmt.Println("-----------------")

	var randomArr [5]int
	len := len(randomArr)
	//set different seed
	rand.Seed(time.Now().UnixNano())
	fmt.Println("different seed")
	for i := 0; i < 5; i++ {
		randomArr[i] = rand.Intn(100)
		//range in 100
		fmt.Println("random number: ", randomArr[i])
	}

	fmt.Println("-----------------")

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if randomArr[j] > randomArr[j+1] {
				randomArr[j], randomArr[j+1] = randomArr[j+1], randomArr[j]
			}
		}
	}
	fmt.Println("after sort:")

	for i := 0; i < len; i++ {
		fmt.Println("sort number: ", randomArr[i])
	}
}
