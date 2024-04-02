package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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
		// initial
		var s []string // do not occupy the ram ... is empty . is nil
		fmt.Printf("var s []string: empty=%t \t nil=%t \n", len(s) == 0, s == nil)

		s = []string{} // is empty , not nil
		fmt.Printf("[]string{}: empty=%t \t nil=%t \n", len(s) == 0, s == nil)

		s = []string(nil) // do not occupy the ram ... is empty . is nil
		fmt.Printf("[]string(nil): empty=%t \t nil=%t \n", len(s) == 0, s == nil)

		s = make([]string, 0) // is empty , not nil
		fmt.Printf("make([]string, 0): empty=%t \t nil=%t \n", len(s) == 0, s == nil)

		fmt.Println()
	}
	{
		// COPY
		//1)
		src := []int{0, 1, 2}
		dst := make([]int, len(src))
		copy(dst, src)
		fmt.Println(dst)
		//2)
		src = []int{0, 1, 2}
		dst = append([]int(nil), src...)

		fmt.Println()
	}
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

		// array
		{
			var numbers [5]int
			fmt.Printf("num array : %#v\n", numbers)
			fmt.Printf("len(nums) : %d\n", len(numbers))
		}
		{
			var nums []int
			fmt.Printf("nums slice: %#v\n", nums)
			fmt.Printf("len(nums) : %d\n", len(nums))
		}
		{
			var array [2]int
			var slice []int
			fmt.Println("slice == nil?", slice == nil)
			fmt.Println("len(slice)  :", len(slice))
			fmt.Printf("array's type: %T\n", array)
			fmt.Printf("slice's type: %T\n", slice)
			rand.Seed(time.Now().UnixNano())
			fmt.Println("--------------")
		}
		{
			// get the unique integer by range
			const max = 5
			var uniques [max]int

		loop:
			for found := 0; found < max; {
				n := rand.Intn(max) + 1
				fmt.Print(n, " ")

				for _, u := range uniques {
					if u == n {
						continue loop
					}
				}

				uniques[found] = n
				found++
			}

			fmt.Println("\n\n uniques:", uniques)
			fmt.Println("--------------")
		}
		{
			rand.Seed(time.Now().UnixNano())
			max := 5
			var uniques []int

		loop2:
			// you can still use the len function on a nil slice
			for len(uniques) < max {
				n := rand.Intn(max) + 1
				fmt.Print(n, " ")

				for _, u := range uniques {
					if u == n {
						continue loop2
					}
				}

				uniques = append(uniques, n)
			}

			fmt.Println("\n\nuniques:", uniques)
			fmt.Println("\nlength of uniques:", len(uniques))

			sort.Ints(uniques)
			fmt.Println("\nsorted:", uniques)

			// convert to slice to use some function
			nums := [5]int{5, 4, 3, 2, 1}
			fmt.Println("\nnums:", nums)
			fmt.Println("\n num[1:2]:", nums[1:2])
			fmt.Println("\n num[1:4]:", nums[0:4])

			sort.Ints(nums[:])

			fmt.Println("--------------")

		}
		{
			agesArray := [3]int{35, 15, 25}
			ages := agesArray[0:3]

			ages[0] = 100
			ages[2] = 50

			fmt.Printf("agesArray type : %T , %[1]v\n", agesArray[:])
			fmt.Printf("agesArray's ages type : %T, %[1]v\n", ages)
		}
		// function
		{
			fmt.Println(make([]int, 3, 5))
			evens := []int{2, 4}
			odds := []int{3, 5, 7}

			fmt.Println("evens [before]", evens)
			fmt.Println("odds  [before]", odds)

			N := copy(evens, odds)
			fmt.Printf("%d element(s) are copied.\n", N)

			fmt.Println("evens [after]", evens)
			fmt.Println("odds  [after]", odds)

		}
		{
			scores := []int{1, 2, 3, 4, 5}
			cloneScores := make([]int, 4)

			copy(cloneScores, scores[:len(scores)-2])
			fmt.Println(cloneScores) // [1,2,3, 0]
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
}
