package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	index := arr[0]

	var leftArr, rightArr []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < index {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	leftArr = quickSort(leftArr)
	rightArr = quickSort(rightArr)

	return append(append(leftArr, index), rightArr...)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomArr []int
	for i := 0; i < 10; i++ {
		randomArr = append(randomArr, rand.Intn(100))
	}
	fmt.Println(randomArr)
	fmt.Println("====After the quick sort=====")
	fmt.Println(quickSort(randomArr))
}
