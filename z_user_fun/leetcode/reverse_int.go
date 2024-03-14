package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
*/
func main() {
	sample := -12345678
	fmt.Println(reverse01(sample))
	fmt.Println(reverse02(sample))

}

func reverse02(x int) int {
	res := 0

	for x != 0 {
		currentLast := x % 10
		x = x / 10
		res = res*10 + currentLast

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	return res
}

func reverse01(x int) int {
	isLessNumber := false
	intToString := strconv.Itoa(x)

	if intToString[0] == '-' {
		isLessNumber = true
		intToString = intToString[1:]
	}

	var result = []rune{}

	for i := len(intToString) - 1; i >= 0; i-- {
		result = append(result, rune(intToString[i]))
	}

	intResult, _ := strconv.Atoi(string(result))

	if isLessNumber {
		intResult = -intResult
	}

	if intResult > math.MaxInt32 || intResult < math.MinInt32 {
		return 0
	}

	return intResult
}
