package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("int64 `%d` to binary: %s\n", int64(13), strconv.FormatInt(int64(13), 2))
	fmt.Println(strings.Repeat("----", 20))

	binstr := "101010100"
	baseTen, _ := strconv.ParseInt(binstr, 2, 0)
	fmt.Printf("bin str `%s` to int64 : %d\n", binstr, baseTen)
	fmt.Println(strings.Repeat("----", 20))

	// custom function to convert int to bin
	var num = 8
	str := func(num int) string {
		var result []int
		var str string

		for ; num > 0; num /= 2 {
			result = append(result, num%2)
		}
		for i := len(result) - 1; i >= 0; i-- {
			str = fmt.Sprintf("%s%d", str, result[i])
		}
		return str
	}(num)
	fmt.Printf("int32 `%d` 2 str bin : %s\n", num, str)
	fmt.Println(strings.Repeat("----", 20))
}
