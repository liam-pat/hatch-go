package main

import (
	"fmt"
	"strings"
)

/*
*
arr := [...]int{10,23,44,55,66}
slice := arr[0:3:5]

[low:high:max]
low: start
high: end
eg. [a[low],a[high])

len：length = high - low
cap：cap = max - low
*/
func main() {
	{
		// initialize a slice
		var s []string // s = nil
		fmt.Printf("`[]string len` len = %d, s == nil is %t, add: %p\n", len(s), s == nil, s)

		s = []string(nil)
		fmt.Printf("`[]string(nil)` len = %d, s == nil is %t, add: %p\n", len(s), s == nil, s)

		// note , []string{} and make([]string, 0) the address is the same
		s = []string{} // s = {}
		fmt.Printf("`[]string{}` len = %d, s == nil is %t, add: %p\n", len(s), s == nil, s)

		s1 := make([]string, 0) // s= {}
		fmt.Printf("`make([]string, 0)` len = %d, s == nil is %t ,add: %p\n", len(s1), s1 == nil, s1)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		src := []int{0, 1, 2}
		dst := make([]int, len(src))

		copy(dst, src)
		fmt.Println("copy src to dst: ", dst)

		src = []int{4, 5, 6}
		dst = append([]int(nil), src...)
		fmt.Println("append scr to dst", dst)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		array := [...]int{10, 23, 44, 55, 66}

		slice := array[1:4:5]
		fmt.Printf("array: %v, array[1:4:5]: %v len: %d cap: %d\n", array, slice, len(slice), cap(slice))
		//if the cap is not enough , it will create automatically,
		slice = append(slice, 100)
		fmt.Printf("append `100` to slice: %v len: %d cap: %d\n", slice, len(slice), cap(slice))

		slice2 := slice[1:]
		slice2[2] = 777
		fmt.Printf("array: %v, slice2: %v, slice: %v\n", array, slice2, slice)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		// auto increament the cap
		sliceS := make([]int, 0, 1)
		ordinalCap := cap(sliceS)

		for i := 0; i < 20; i++ {
			sliceS = append(sliceS, i)

			if newCap := cap(sliceS); ordinalCap < newCap {
				fmt.Printf("original cap : %3d ,new cap: %3d \n", ordinalCap, newCap)
				ordinalCap = newCap
			}
		}

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		scr := []int{1, 2}
		dst := []int{6, 6, 6, 6, 6, 6}
		copy(dst, scr)
		fmt.Printf("scr: %v, dst: %v\n", scr, dst)
		fmt.Println("copy scr to dst: ", dst)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		arr := [3]int{35, 15, 25}
		ages := arr[0:3]

		ages[0] = 100
		ages[2] = 50

		fmt.Printf("arr type : %T , %[1]v\n", arr[:])
		fmt.Printf("arr's ages type : %T, %[1]v\n", ages)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		evens := []int{2, 4}
		odds := []int{3, 5, 7}
		fmt.Printf("evens: %v, odds: %v\n", evens, odds)

		N := copy(evens, odds)
		fmt.Printf("copy odds to evens, %d element(s) are copied. result: %v\n", N, evens)

		fmt.Println(strings.Repeat("###", 20))
	}
	{
		scr := []int{1, 2, 3, 4, 5}
		dst := make([]int, 4)
		fmt.Printf("scr: %v, dst: %v\n", scr, dst)

		copy(dst, scr[:len(scr)-2])
		fmt.Printf("copy `scr[:len(scr)-2]` to dst: %v\n", dst) // [1,2,3, 0]
	}
}
