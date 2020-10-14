package main

import "fmt"

func lengthOfNonRepeatingSubString(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if existIndex, ok := lastOccurred[ch]; ok && existIndex >= start {
			start = existIndex + 1
		}
		// 当前位置减去开始的位置再加上 1.若比 maxLength，那么字符串就有另外一个比第一个长，
		if i-start+1 >= maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubString("abcdabde"))
}
