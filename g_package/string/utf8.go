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
		// 1 漢字 use 3 bytes , 24 bit, meansa [3]byte
		name := "大中國"

		fmt.Printf("%s len = %d \n", name, len(name))

		fmt.Printf("[0]byte = %d, %[1]b \n", name[0])
		fmt.Printf("[1]byte = %d, %[1]b \n", name[1])
		fmt.Printf("[2]byte = %d, %[1]b \n", name[2])
		fmt.Printf("[3]byte = %d, %[1]b \n", name[3])
		fmt.Printf("[4]byte = %d, %[1]b \n", name[4])
		fmt.Printf("[5]byte = %d, %[1]b \n", name[5])
		fmt.Printf("[6]byte = %d, %[1]b \n", name[6])
		fmt.Printf("[7]byte = %d, %[1]b \n", name[7])
		// strconv.FormatInt sample
		fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

		byteSet := []byte(name)
		fmt.Printf("[]byte : %+v \n", byteSet)

		runeSet := []rune(name)
		fmt.Printf("[]rune : %+v \t len: %d \n", runeSet, utf8.RuneCountInString(name))

		Runes, size := utf8.DecodeRuneInString("測")
		fmt.Printf("str: %s \nrune : %d \nbyte size : %d \n", string(Runes), Runes, size)

		fmt.Println(strings.Repeat("##", 10))
	}
	{
		word := []byte("界")

		fmt.Printf("utf8.Valid(word[:2]) = %t \n", utf8.Valid(word[:2]))
		fmt.Printf("utf8.ValidRune('界') = %t \n", utf8.ValidRune('界'))
		fmt.Printf("utf8.ValidString(\"世界\") = %t \n", utf8.ValidString("世界"))

		fmt.Printf("utf8.RuneLen('界') = %d \n", utf8.RuneLen('界'))
		fmt.Printf("utf8.RuneCount(word) = %d \n", utf8.RuneCount(word))
		fmt.Printf("utf8.RuneCountInString(\"世界\") = %d \n", utf8.RuneCountInString("世界"))

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

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		words := []rune{'𝓐', '𝓑'}

		fmt.Printf("utf16.Encode(words) = %v \n", utf16.Encode(words))
		fmt.Printf("utf16.Decode(utf16.Encode(words)) = %v \n", utf16.Decode(utf16.Encode(words)))

		r1, r2 := utf16.EncodeRune('𝓐')
		fmt.Printf("r1(utf16.EncodeRune('𝓐')) = %d, r2(utf16.EncodeRune('𝓐')) = %d \n", r1, r2)
		fmt.Printf("utf16.DecodeRune(r1, r2) = %d \n", utf16.DecodeRune(r1, r2))
		fmt.Printf("utf16.IsSurrogate(r1) = %t \n", utf16.IsSurrogate(r1))
		fmt.Printf("utf16.IsSurrogate(r2) = %t \n", utf16.IsSurrogate(r2))
		fmt.Printf("utf16.IsSurrogate(1234) = %t \n", utf16.IsSurrogate(1234))

		fmt.Println(strings.Repeat("##", 20))
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
		fmt.Println()
	}
}
