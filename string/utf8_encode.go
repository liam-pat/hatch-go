package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	name := "朱晓琛"

	// one chinese use three byte , 24 bit

	// save 毕 in utf-8
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	// save 永 in utf-8
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))

	// save 垚 in utf-8
	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	// echo 9 . 24 bit save one chinese
	fmt.Println(len(name))

	// to slice byte
	byteSet := []byte(name)
	fmt.Println(byteSet)

	// to slice rune (unicode) , int32 , 32 bit save one chinese
	tempSet := []rune(name)
	fmt.Println(tempSet)

	runeLen := utf8.RuneCountInString(name)
	fmt.Println(runeLen)

	cha1 := string(65)
	fmt.Println(cha1)

	cha2 := string(26417)
	fmt.Println(cha2)

	v3, size := utf8.DecodeRuneInString("毕")
	fmt.Println(v3, size)

}
