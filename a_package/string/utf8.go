package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	{
		name := "大中國"

		fmt.Println("len is : ", len(name))

		// 1 Chinese char use 3 bytes , 24 bit
		fmt.Println("INT | BINARY")
		fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
		fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
		fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))
		fmt.Println(strings.Repeat("大-", 5))

		fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
		fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
		fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))
		fmt.Println(strings.Repeat("中-", 5))

		fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
		fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
		fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))
		fmt.Println(strings.Repeat("國-", 5))

		fmt.Println(strings.Repeat("#*", 10))

		byteSet := []byte(name)
		fmt.Printf("convert to []byte : %v\n", byteSet)
		fmt.Println(strings.Repeat("#*", 10))

		tempSet := []rune(name)
		fmt.Printf("convert to []rune : %v \t rune len : %d\n", tempSet, utf8.RuneCountInString(name))
		fmt.Println(strings.Repeat("#*", 10))

		runeTmp, byteSize := utf8.DecodeRuneInString("毕")
		cha2 := string(runeTmp)

		fmt.Println(cha2, runeTmp, byteSize)
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		word := []byte("界")

		fmt.Println(utf8.Valid(word[:2]))
		fmt.Println(utf8.ValidRune('界'))
		fmt.Println(utf8.ValidString("世界"))
		fmt.Println(strings.Repeat("#*", 10))

		fmt.Println(utf8.RuneLen('界'))
		fmt.Println(utf8.RuneCount(word))
		fmt.Println(utf8.RuneCountInString("世界"))
		fmt.Println(strings.Repeat("#*", 10))

		p := make([]byte, 3)
		utf8.EncodeRune(p, '好')
		fmt.Println(p)
		fmt.Println(utf8.DecodeRune(p))
		fmt.Println(utf8.DecodeRuneInString("你好"))
		fmt.Println(utf8.DecodeLastRune([]byte("你好")))
		fmt.Println(utf8.DecodeLastRuneInString("你好"))
		fmt.Println(strings.Repeat("#*", 10))

		fmt.Println(utf8.FullRune(word[:2]))
		fmt.Println(utf8.FullRuneInString("你好"))
		fmt.Println(strings.Repeat("#*", 10))

		fmt.Println(utf8.RuneStart(word[1]))
		fmt.Println(utf8.RuneStart(word[0]))
	}
}
