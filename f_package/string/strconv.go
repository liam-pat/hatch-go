package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	{
		// append to slice
		slice := make([]byte, 0, 1024)
		slice = strconv.AppendBool(slice, true)
		slice = strconv.AppendInt(slice, 1234, 10)
		slice = strconv.AppendQuote(slice, "test")
		slice = strconv.AppendQuoteRune(slice, '单')
		fmt.Println("byte slice: ", slice)
		fmt.Println("string slice: ", string(slice))

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		// format .other type convert to string
		fmt.Printf("type: %T, value : %[1]v\n", strconv.FormatBool(true))
		fmt.Printf("type: %T, value : %[1]v\n", strconv.FormatFloat(3.14, 'f', -1, 64))
		fmt.Printf("type: %T, value : %[1]v\n", strconv.Itoa(999))
		fmt.Printf("type: %T, value : %[1]v\n", strconv.FormatInt(1234, 10))
		fmt.Printf("type: %T, value : %[1]v\n", strconv.FormatUint(12345, 10))

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Println(strconv.ParseBool("false"))
		fmt.Println(strconv.ParseFloat("123.23", 64))
		fmt.Println(strconv.ParseInt("1234", 10, 64))
		fmt.Println(strconv.ParseUint("12345", 10, 64))
		fmt.Println(strconv.Atoi("1023"))
		fmt.Println(strings.Repeat("##", 20))
	}
}
