package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	{
		// quote by '' , it would be the unicode/rune string
		fmt.Printf("%-50s = %t\n", "unicode.IsControl('\\u0015')", unicode.IsControl('\u0015'))
		fmt.Printf("%-50s = %t\n", "unicode.IsControl('\\ufe35')", unicode.IsControl('\ufe35'))

		fmt.Printf("%-50s = %t\n", "unicode.IsDigit('1')", unicode.IsDigit('1'))
		fmt.Printf("%-50s = %t\n", "unicode.IsNumber('1')", unicode.IsNumber('1'))

		fmt.Printf("%-50s = %t\n", "unicode.IsDigit('Ⅷ')", unicode.IsDigit('Ⅷ'))
		fmt.Printf("%-50s = %t\n", "unicode.IsNumber('Ⅷ')", unicode.IsNumber('Ⅷ'))

		fmt.Printf("%-50s = %t\n", "unicode.IsDigit('你')", unicode.IsDigit('你'))
		fmt.Printf("%-50s = %t\n", "unicode.Is(unicode.Han, '你')", unicode.Is(unicode.Han, '你'))
		fmt.Printf("%-50s = %t\n", "unicode.In('你', unicode.Gujarati, unicode.White_Space)", unicode.In('你', unicode.Gujarati, unicode.White_Space))
		fmt.Println(strings.Repeat("##", 10))
	}
}
