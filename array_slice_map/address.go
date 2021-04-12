package main

import "fmt"

func main() {

	/**
	array_slice_map will copy one
	slice will point to the same address which save the value
	*/
	nums := [3]int8{11, 22, 33}

	demo := nums
	fmt.Printf("nums address : %p \n", &nums)
	fmt.Printf("demo address : %p \n", &demo)

	fmt.Printf("array_slice_map address : %p \n", &nums)
	fmt.Printf("array_slice_map[1] address : %p \n", &nums[0])
	fmt.Printf("array_slice_map[2] address : %p \n", &nums[1])
	fmt.Printf("array_slice_map[3] address : %p \n", &nums[2])
}
