package main

import "fmt"

/**
array := [...]int{10,23,44,55,66}
slice := array[0:3:5]

[low:high:max]
low: start
high: end
eg. [a[low],a[high])  左闭右开

长度：length = high - low
容量：cap = max = low

切片 s[i] 不可超越 len(s) 向后拓展 不可以超越底层数组 cap（s）
append 之后会新增一个数组，之前的那个数组会被回收掉
 */

func main() {
	array := [...]int{10, 23, 44, 55, 66}
	slice := array[1:4:5]

	fmt.Println("slice: ", slice)
	fmt.Println("slice length: ", len(slice))
	fmt.Println("slice cap: ", cap(slice))

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	fmt.Println("---------------")

	//you can user append fun to append an element,if the cap is not enough , if will create automatically,
	slice = append(slice, 1)
	fmt.Println("after use append function , slice: ", slice)
	fmt.Println("---------------")

	// if you the slice alter ,slice2 also alter,because it use the address to tran
	fmt.Println("after changing slice2 , the slice also change")
	slice2 := slice[1:]
	slice2[2] = 777
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice: ", slice)
	fmt.Println("---------------")

	//会自动扩容，按两倍的速度
	fmt.Println("test the cap's growth ")
	s := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, 1)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("old cap :%d  new cap: %d \n", oldCap, newCap)
			oldCap = newCap
		}
	}
	fmt.Println("---------------")

	//test the copy function
	srcSlice := []int{1, 2}
	disSlice := []int{6, 6, 6, 6, 6, 6}
	copy(disSlice, srcSlice)
	fmt.Println("the src slice copy to dis slice result : ",disSlice)
	fmt.Println("---------------")

	//you can init slice in two ways
	//1. by auto : s : = []int{1,2,3,4}
	//2. by make : s := make([]int ,length , cap)   if cap is null,cap also equal length
}

/**
different from array to slice
array : length cap always const
slice :  [] is null or ... ,length and cap can change
 */
