package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRandNum() int {
	var number int
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		number = rand.Intn(10000)
		if number > 1000 {
			break
		}
	}
	return number
}

func num2slice(num int) []int {
	slice := make([]int, 4)
	slice[0] = num / 1000
	slice[1] = num % 1000 / 100
	slice[2] = num % 1000 % 100 / 10
	slice[3] = num % 10
	return slice
}

func onGame(randNum int) {
	slice := num2slice(randNum)
	for {
		var input int
		var equalNum int = 0
		for {
			fmt.Printf("input number:")
			fmt.Scan(&input)
			if input >= 10000 || input < 1000 {
				fmt.Println("<<<<< notice the rule (1000 < num < 10000)")
				continue
			}
			break
		}
		inputNumToSlice := num2slice(input)
		for key, value := range slice {
			if value > inputNumToSlice[key] {
				fmt.Println(fmt.Sprintf("digital: %d) is small", inputNumToSlice[key]))
			} else if value < inputNumToSlice[key] {
				fmt.Println(fmt.Sprintf("digital: %d) is bigger", inputNumToSlice[key]))
			} else {
				fmt.Println(fmt.Sprintf("digital: %d) is right", inputNumToSlice[key]))
				equalNum++
			}
		}
		if equalNum == len(slice) {
			fmt.Println("bingo")
			break
		}
	}

}

func main() {
	randNum := createRandNum()
	onGame(randNum)
}
