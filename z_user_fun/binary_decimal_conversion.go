package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// decimal to binary-string
	var number int64 = 13
	binaryStr := strconv.FormatInt(number, 2)
	fmt.Printf("int64 `%d` to binary string : %s\n", number, binaryStr)
	fmt.Println(strings.Repeat("----", 20))

	// binary-string to int64
	intStr := "101010100"
	baseTen, _ := strconv.ParseInt(intStr, 2, 0)
	fmt.Printf("binary string `%s` convert to int64 : %d\n", intStr, baseTen)
	fmt.Println(strings.Repeat("----", 20))

	// custom function to convert int to bin
	var number32 = 8
	str := func(number32 int) string {
		var result []int
		for ; number32 > 0; number32 /= 2 {
			result = append(result, number32%2)
		}
		var toStr string
		for i := len(result) - 1; i >= 0; i-- {
			toStr = fmt.Sprintf("%s%d", toStr, result[i])
		}
		return toStr
	}(number32)
	fmt.Printf("int32 `%d` convert to string binary : %s\n", number32, str)
	fmt.Println(strings.Repeat("----", 20))
}
