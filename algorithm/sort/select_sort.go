package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
选择排序是一种不稳定的排序，需要记录位置，
*/
func selectSort(arr []int) []int {
	length := len(arr)
	//从第一个元素开始
	for i := 0; i < length; i++ {
		position := i
		//从标志位的后一位开始选择
		for j := i + 1; j < length; j++ {
			//有比标志位小的记住位置，直到找到最小的
			if arr[j] < arr[position] {
				position = j
			}
		}
		if position != i {
			//找到这一轮最小的跟标志位对换
			arr[position] = arr[i] + arr[position]
			arr[i] = arr[position] - arr[i]
			arr[position] = arr[position] - arr[i]
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
	fmt.Println("====After the select sort=====")
	fmt.Println(selectSort(randomArr))
}
