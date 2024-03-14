package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 插入排序
func insertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		if arr[i] < arr[i-1] {
			j := 0                                       //more wide
			tag := arr[i]                                //sign
			for j = i - 1; j >= 0 && arr[j] > tag; j-- { // members back to move
				arr[j+1] = arr[j]
			}
			arr[j+1] = tag //insert
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
	fmt.Println(randomArr)
	fmt.Println("====After the insert sort=====")
	fmt.Println(insertSort(randomArr))
}
