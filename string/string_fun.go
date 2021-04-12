package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// check whether has goddess in it , has return ture , not  return false
	fmt.Println(strings.Contains("hello my goddess", "goddess"))

	// implode slice to string
	s := []string{"package", "joy"}
	fmt.Println(strings.Join(s, "-"))

	// index ,return the index in the string , not return -1
	fmt.Println(strings.Index("biyongyao", "bi"))

	// repeat string
	fmt.Println(strings.Repeat("love you", 2))

	// split string to array_slice_map
	str := "biyongyao.yy.bi"
	fmt.Println(strings.Split(str, "."))

	// trim two side sign
	fmt.Println(strings.Trim("   {name:'biyongyao'} ", " "))

	// explode to a slice by space
	fmt.Println(strings.Fields(" name sex key   "))

	// convert to string save []byte
	fmt.Println("---------")
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abdfsd")
	fmt.Println(string(slice))

	//other type convert to string
	fmt.Println("---------")
	var str2 string
	str2 = strconv.FormatBool(true)
	str2 = strconv.FormatFloat(3.14, 'f', -1, 64)
	str2 = strconv.Itoa(999)
	fmt.Println(str2)

	// string convert to other type
	fmt.Println("---------")
	flag, _ := strconv.ParseBool("true")
	fmt.Printf("flag type = %T,flag value = %t ", flag, flag)

	// string convert to int
	fmt.Println("---------")
	ia, _ := strconv.Atoi("1234")
	fmt.Printf("ia type = %T,ia value = %d ", ia, ia)

}
