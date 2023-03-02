package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	{
		str := "hey"
		bytes := []byte{104, 101, 121}
		// same as: []byte("hey")
		fmt.Printf(`"hey" as bytes   : %d`+"\n", []byte(str))
		// same as: string([]byte{104, 101, 121})
		fmt.Printf("bytes as string  : %q\n", string(bytes))

		fmt.Println(strings.Repeat("**##", 20))
		fmt.Printf("%c                : %[1]d\n", 'h')
		fmt.Printf("%c                : %[1]d\n", 'e')
		fmt.Printf("%c                : %[1]d\n", 'y')

		fmt.Println(strings.Repeat("**##", 20))
		var (
			anInt   int   = 'h'
			anInt8  int8  = 'h'
			anInt16 int16 = 'h'
			anInt32 int32 = 'h'
			// rune literal's default type is: rune
			aRune = 'h'
		)
		fmt.Printf("rune literals are typeless:\t %T %T %T %T %T\n", anInt, anInt8, anInt16, anInt32, aRune)
		fmt.Println(strings.Repeat("**##", 10))
		fmt.Printf("%q in decimal: %[1]d\n", 104)
		fmt.Printf("%q in binary : %08[1]b\n", 'h')
		fmt.Printf("%q in hex    : 0x%[1]x\n", 0x68)
	}
	{
		fmt.Println(strings.Repeat("**##", 20))
		str := "Yūgen ☯ 💀"
		bytes := []byte(str)
		str = string(bytes)

		fmt.Printf("str output : %s\n", str)
		fmt.Printf("\t%d bytes\n", len(str))
		fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

		fmt.Printf("rune output : % x\n", bytes)
		fmt.Printf("\t%d bytes\n", len(bytes))
		fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

		fmt.Println(strings.Repeat("**##", 20))
		fmt.Printf("1st byte   : %c\n", str[0])           // ok
		fmt.Printf("2nd byte   : %c\n", str[1])           // not ok
		fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
		fmt.Printf("last rune  : %s\n", str[len(str)-4:]) // ok
	}
	{
		// string , rune comparison
		fmt.Println(strings.Repeat("**##", 20))

		str := "Yūgen ☯ 💀"
		runes := []rune(str)
		fmt.Printf("string output : %s\n", str)
		fmt.Printf("\t runes len : %d \n", len(runes))
		fmt.Printf("\t runes[] bytes : %d \n", int(unsafe.Sizeof(runes[0]))*len(runes))

		fmt.Printf("1st rune   : %c\n", runes[0])
		fmt.Printf("2nd rune   : %c\n", runes[1])
		fmt.Printf("runes[:5] : %c\n", runes[:5])
	}
	{
		// rune ,byte comparison
		fmt.Println(strings.Repeat("**##", 20))
		word := "öykü"
		fmt.Printf("%q in runes: %c ** rune len = %d ** %d bytes\n", word, []rune(word), len([]rune(word)), int(unsafe.Sizeof([]rune(word)[0]))*len([]rune(word)))
		fmt.Printf("%q in bytes: % [1]x ** byte len = %d ** %d bytes\n", word, len([]byte(word)), int(unsafe.Sizeof([]byte(word)[0]))*len([]byte(word)))

		fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]}) // ö
		fmt.Printf("%c\n", word[2])                               // y
		fmt.Printf("%c\n", word[3])                               // k
		fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
	}
	{
		// rune decoding
		fmt.Println(strings.Repeat("**##", 20))

		r, size := utf8.DecodeRuneInString("öykü")
		fmt.Printf("rune: %c size: %d bytes.\n", r, size)

		r, size = utf8.DecodeRuneInString("ykü")
		fmt.Printf("rune: %c size: %d bytes.\n", r, size)

		r, size = utf8.DecodeRuneInString("kü")
		fmt.Printf("rune: %c size: %d bytes.\n", r, size)

		r, size = utf8.DecodeRuneInString("ü")
		fmt.Printf("rune: %c size: %d bytes.\n", r, size)

		const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede.`
		for i := 0; i < len(text); {
			r, size := utf8.DecodeRuneInString(text[i:])
			fmt.Printf("%c", r)
			i += size
		}
	}
}
