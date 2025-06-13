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
		fmt.Printf("string to byte  : %d\n", []byte(str))
		fmt.Printf("bytes 2 string  : %q\n", string(bytes))

		fmt.Printf("%c : %[1]d\n", 'h')
		fmt.Printf("%c : %[1]d\n", 'e')
		fmt.Printf("%c : %[1]d\n", 'y')

		fmt.Printf("%q in decimal: %[1]d\n", 'h')
		fmt.Printf("%q in binary : %0[1]b\n", 'h')
		fmt.Printf("%q in hex    : 0x%[1]x\n", 0x68)
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		str := "Yūgen ☯ 💀"
		bytes := []byte(str)

		fmt.Printf("str: %s\n", str)
		fmt.Printf("%8d bytes\n", len(str))
		fmt.Printf("%8d runes\n", utf8.RuneCountInString(str))

		fmt.Printf("rune: % x\n", bytes)
		fmt.Printf("%8d bytes\n", len(bytes))
		fmt.Printf("%8d runes\n", utf8.RuneCount(bytes))

		fmt.Printf("1st byte   : %c\n", str[0]) // ok
		fmt.Printf("2nd byte   : %c\n", str[1]) // not ok

		fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
		fmt.Printf("last rune  : %s\n", str[len(str)-4:]) // ok
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		str := "Yūgen ☯ 💀"
		runes := []rune(str)

		fmt.Printf("str: %s\n", str)
		fmt.Printf("runes len: %d \n", len(runes))
		fmt.Printf("bytes len: %d \n", int(unsafe.Sizeof(runes[0]))*len(runes))

		fmt.Printf("1st: %c\n", runes[0])
		fmt.Printf("2nd: %c\n", runes[1])
		fmt.Printf("runes[:5]: %c\n", runes[:5])
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		word := "öykü"
		fmt.Printf("%q 2 runes: %c \n len = %d \n %d bytes\n", word, []rune(word), len([]rune(word)), int(unsafe.Sizeof([]rune(word)[0]))*len([]rune(word)))
		fmt.Printf("%q 2 bytes: % #[1]X \n len = %d \n %d bytes\n", word, len([]byte(word)), int(unsafe.Sizeof([]byte(word)[0]))*len([]byte(word)))
		fmt.Println(strings.Repeat("###", 20))
	}
	{
		// only decode the first rune
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
