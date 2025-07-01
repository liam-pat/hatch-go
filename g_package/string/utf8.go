package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	{
		// 1 Chinese char use 3 bytes , 24 bit

		name := "大中國"

		fmt.Println("len is : ", len(name))

		fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
		fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
		fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))
		fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
		fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
		fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))
		fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
		fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
		fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

		byteSet := []byte(name)
		fmt.Printf("convert to []byte : %+v ## []byte len : %d \n", byteSet, len(byteSet))

		runeSet := []rune(name)
		fmt.Printf("convert to []rune : %+v ## []rune len : %d \n", runeSet, utf8.RuneCountInString(name))

		runeTmp, size := utf8.DecodeRuneInString("毕")

		fmt.Printf("string: %s, rune : %d, byte size : %d \n", string(runeTmp), runeTmp, size)
		fmt.Println(strings.Repeat("##", 10))
	}
	{
		// UTF8 package
		word := []byte("界")

		fmt.Println(utf8.Valid(word[:2]))
		fmt.Println(utf8.ValidRune('界'))
		fmt.Println(utf8.ValidString("世界"))

		fmt.Println(utf8.RuneLen('界'))
		fmt.Println(utf8.RuneCount(word))
		fmt.Println(utf8.RuneCountInString("世界"))

		p := make([]byte, 3)
		utf8.EncodeRune(p, '好')
		fmt.Println(p)
		fmt.Println(utf8.DecodeRune(p))
		fmt.Println(utf8.DecodeRuneInString("你好"))
		fmt.Println(utf8.DecodeLastRune([]byte("你好")))
		fmt.Println(utf8.DecodeLastRuneInString("你好"))

		fmt.Println(utf8.FullRune(word[:2]))
		fmt.Println(utf8.FullRuneInString("你好"))

		fmt.Println(utf8.RuneStart(word[1]))
		fmt.Println(utf8.RuneStart(word[0]))

		fmt.Println(strings.Repeat("#", 20))
	}
	{
		// UTF16 package
		words := []rune{'𝓐', '𝓑'}

		u16 := utf16.Encode(words)
		fmt.Println(u16)
		fmt.Println(utf16.Decode(u16))

		r1, r2 := utf16.EncodeRune('𝓐')
		fmt.Println(r1, r2)
		fmt.Println(utf16.DecodeRune(r1, r2))
		fmt.Println(utf16.IsSurrogate(r1))
		fmt.Println(utf16.IsSurrogate(r2))
		fmt.Println(utf16.IsSurrogate(1234))

		fmt.Println(strings.Repeat("#", 20))
	}

	{
		b := []byte("你好，世界")
		for k, v := range b {
			fmt.Printf("%d->%s |", k, string(v))
		}
		fmt.Println()
		r := bytes.Runes(b)
		for k, v := range r {
			fmt.Printf("%d->%s|", k, string(v))
		}
		fmt.Println(strings.Repeat("#", 20))
	}
}
