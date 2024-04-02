package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func bubbleSortOptimization01(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		isSort := true
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				// because this time still exchange
				isSort = false
			}
		}
		if isSort == true {
			break
		}
	}
	return arr
}

func bubbleSortOptimization02(arr []int) []int {
	length := len(arr)
	lastExchangeIndex := 0
	exchangeBorder := length - 1
	for i := 0; i < length-1; i++ {
		isSort := true

		for j := 0; j < exchangeBorder; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				// because this time still exchange
				isSort = false

				// mark the last exchange index
				lastExchangeIndex = j
			}
		}
		// next time sort to this index
		exchangeBorder = lastExchangeIndex

		if isSort == true {
			break
		}
	}
	return arr
}

func bubbleSortOptimization03(arr []int) []int {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		isSort := true
		for j := i; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				isSort = false
			}
		}
		if isSort == true {
			break
		}

		isSort = true
		for j := i; j < length-1-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isSort = false
			}
		}
		if isSort == true {
			break
		}
	}

	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomArr []int
	for i := 0; i < 10; i++ {
		randomArr = append(randomArr, rand.Intn(100))
	}
	fmt.Println("==== generate Array ====")
	fmt.Println(randomArr)

	fmt.Println("==== After the common bubble sort =====")
	fmt.Println(bubbleSort(randomArr))

	fmt.Println("==== After the optimization 01 bubble sort =====")
	fmt.Println(bubbleSortOptimization01(randomArr))

	fmt.Println("==== After the optimization 02 bubble sort =====")
	fmt.Println(bubbleSortOptimization02(randomArr))

	fmt.Println("==== After the optimization 03 bubble sort =====")
	fmt.Println(bubbleSortOptimization03(randomArr))
}
