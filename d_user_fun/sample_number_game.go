package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandNum() int {
	rand.Seed(time.Now().UnixNano())

	var number int
	for {
		number = rand.Intn(10000)
		if number > 1000 {
			break
		}
	}

	return number
}

func explodeNumberToSlice(number int) []int {
	returnSlice := make([]int, 4)

	returnSlice[0] = number / 1000
	returnSlice[1] = number % 1000 / 100
	returnSlice[2] = number % 1000 % 100 / 10
	returnSlice[3] = number % 10

	return returnSlice
}

func onGame(number int) {

	numberToSlice := explodeNumberToSlice(number)
	for {
		var input int
		var equalNum int

		for {
			equalNum = 0
			fmt.Printf("input number :")
			fmt.Scan(&input)

			if input >= 10000 || input < 1000 {
				fmt.Println("you input a error number , please try again")
				continue
			}
			break
		}

		inputNumToSlice := explodeNumberToSlice(input)

		for key, value := range numberToSlice {
			if value > inputNumToSlice[key] {
				fmt.Println(fmt.Sprintf("key : %d is small", key))
			} else if value < inputNumToSlice[key] {
				fmt.Println(fmt.Sprintf("key : %d is bigger", key))
			} else {
				fmt.Println(fmt.Sprintf("key : %d is equl", key))
				equalNum++
			}
		}

		if equalNum == len(numberToSlice) {
			fmt.Println("#### All number you guess right ,exit")
			break
		}
	}

}

func main() {
	num := createRandNum()
	onGame(num)
}
