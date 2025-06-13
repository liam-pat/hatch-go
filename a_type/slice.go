package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
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
cap：cap = max - low
*/

func main() {
	{
		var s []string // s = nil
		fmt.Printf("[]string len = %d, s == nil is  %t \n", len(s) == 0, s == nil)

		s = []string(nil) // do not occupy the ram ... is empty . is nil
		fmt.Printf("[]string len = %d, s == nil is  %t \n", len(s) == 0, s == nil)

		s = []string{} // s = {} , not nil
		fmt.Printf("[]string len = %d, s == nil is  %t \n", len(s) == 0, s == nil)

		s = make([]string, 0) // s= {} , not nil
		fmt.Printf("[]string len = %d, s == nil is  %t \n", len(s) == 0, s == nil)
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		src := []int{0, 1, 2}
		dst := make([]int, len(src))

		copy(dst, src)
		fmt.Println(dst)

		src = []int{4, 5, 6}
		dst = append([]int(nil), src...)
		fmt.Println(dst)
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		array := [...]int{10, 23, 44, 55, 66}

		slice := array[1:4:5]
		fmt.Printf("slice : %v len: %d cap: %d\n", slice, len(slice), cap(slice))
		//if the cap is not enough , it will create automatically,
		slice = append(slice, 100)
		fmt.Println("after appending slice: ", slice)
		fmt.Println("original array: ", array)

		slice2 := slice[1:]
		slice2[2] = 777
		fmt.Println("slice2: ", slice2)
		fmt.Println("slice: ", slice)
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		array := make([]int, 0, 1) // other way array = []int{}
		oldCap := cap(array)
		for i := 0; i < 20; i++ {
			array = append(array, 1)
			if newCap := cap(array); oldCap < newCap {
				fmt.Printf("old cap : %d  \t new cap: %d \n", oldCap, newCap)
				oldCap = newCap
			}
		}
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		scr := []int{1, 2}
		dst := []int{6, 6, 6, 6, 6, 6}
		copy(dst, scr)
		fmt.Println("result: ", slice02)
		fmt.Println(strings.Repeat("###", 20))
	}
	{
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
	}
}
