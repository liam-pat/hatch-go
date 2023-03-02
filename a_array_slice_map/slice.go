package main

import "fmt"

/**
arr := [...]int{10,23,44,55,66}
slice := arr[0:3:5]

[low:high:max]
low: start
high: end
eg. [a[low],a[high])

len：length = high - low
cap：cap = max = low

切片 s[i] 不可超越 len(s) 向后拓展 不可以超越底层数组 cap（s）
append function will remove pre slice and then create new one
*/

func main() {
	{
		array := [...]int{10, 23, 44, 55, 66}
		slice := array[1:4:5]

		fmt.Println("slice : ", slice)
		fmt.Println("slice length : ", len(slice))
		fmt.Println("slice cap : ", cap(slice))

		//if the cap is not enough , it will create automatically,
		slice = append(slice, 1)
		fmt.Println("after appending , slice: ", slice)

		slice2 := slice[1:]
		slice2[2] = 777
		fmt.Println("slice2: ", slice2)
		fmt.Println("slice: ", slice)
	}
	{
		array := make([]int, 0, 1) // other way array = []int{}
		originalCap := cap(array)

		for i := 0; i < 20; i++ {
			array = append(array, 1)
			if newCap := cap(array); originalCap < newCap {
				fmt.Printf("original cap : %d  \t new cap: %d \n", originalCap, newCap)
				originalCap = newCap
			}
		}
	}
	{
		slice01 := []int{1, 2}
		slice02 := []int{6, 6, 6, 6, 6, 6}
		copy(slice02, slice01)
		fmt.Println("copy result : ", slice02)
	}
	{
		slice := make([]int, 3)
		// clone new one into fun
		TestSlicePass(slice)
		fmt.Printf("main p %p \n", slice) // still the s1 , cap and len also equal 3
		fmt.Println(slice)                // [ 0 0 0]
	}

	/**
	output:
	slice :  [23 44 55]
	slice length :  3
	slice cap :  4
	after appending , slice:  [23 44 55 1]
	slice2:  [44 55 777]
	slice:  [23 44 55 777]
	original cap : 1  	 new cap: 2
	original cap : 2  	 new cap: 4
	original cap : 4  	 new cap: 8
	original cap : 8  	 new cap: 16
	original cap : 16  	 new cap: 32
	copy result :  [1 2 6 6 6 6]
	fun p 0x1400012e000
	fun p 0x140001240c0
	[0 0 0 1 2 3 4 5]
	main p 0x1400012e000
	[0 0 0]
	*/
}

func TestSlicePass(slice []int) {
	fmt.Printf("fun p %p \n", slice) // as same as main
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("fun p %p \n", slice) // over , it will add the cap auto
	fmt.Println(slice)               // [0 0 0 1 2 3 4 5]
}
