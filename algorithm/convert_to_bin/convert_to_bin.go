package main

import (
	"fmt"
)

func convertToBin(number int) string {
	var result []int
	for ; number > 0; number /= 2 {
		result = append(result, number%2)
	}

	toStr := ""
	for i := len(result) - 1; i >= 0; i-- {
		toStr = fmt.Sprintf("%s%d", toStr, result[i])
	}

	return toStr
}

func main() {
	fmt.Println(convertToBin(9), convertToBin(11))
}
