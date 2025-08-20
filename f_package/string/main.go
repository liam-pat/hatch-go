package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	{
		//comparison
		fmt.Println(strings.Compare("go", "PHP"))
		fmt.Println(strings.EqualFold("GO", "go"))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// substring
		fmt.Println(strings.Contains("golang", "go"))
		fmt.Println(strings.ContainsAny("golang", "g a"))
		fmt.Println(strings.ContainsRune("中國", rune('中')))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// count
		fmt.Println(strings.Count("cheese", "e"))
		fmt.Println(len("谷歌中国"))
		fmt.Println(strings.Count("谷歌中国", ""))

		fmt.Println(strings.Repeat("#*", 10))
	}

	{
		//split
		fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
		fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
		fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
		fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ",")) // ["foo," "bar," "baz"]
		fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))  // ["foo" "bar,baz"]

		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// prefix or suffix
		fmt.Println(strings.HasPrefix("Gopher", "Go"))
		fmt.Println(strings.HasPrefix("Gopher", ""))
		fmt.Println(strings.HasSuffix("Amigo.txt", "txt"))
		fmt.Println(strings.HasSuffix("Amigo", ""))

		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		//index
		fmt.Println(strings.Index("Gopher", "ph"))
		fmt.Println(strings.IndexByte("Gopher", byte('e')))
		fmt.Println(strings.IndexAny("Gopher", "a e"))
		fmt.Println(strings.IndexRune("中國", rune('國')))

		fmt.Println(strings.IndexFunc("Hello, 世界", func(r rune) bool {
			return unicode.Is(unicode.Han, r)
		}))

		fmt.Println(strings.LastIndex("Gopher Gopher", "ph"))
		fmt.Println(strings.LastIndexByte("Gopher Gopher", byte('e')))
		fmt.Println(strings.LastIndexAny("Gopher Gopher", "a e"))

		fmt.Println(strings.LastIndexFunc("Hello, 世界 >> Hello, 世界", func(r rune) bool {
			return unicode.Is(unicode.Han, r)
		}))

		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// string join
		fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// string repeat
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		//mapping
		mapping := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return r + 32
			case r >= 'a' && r <= 'z':
				return r
			case unicode.Is(unicode.Han, r):
				return '\n'
			}
			return -1 // 过滤所有非字母、汉字的字符
		}
		fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))
		fmt.Println(strings.Repeat("#*", 10))
	}

	{
		// string replace
		fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
		fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
		fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		fmt.Println(strings.ToLower("HELLO WORLD"))
		fmt.Println(strings.ToLower("Ā Á Ǎ À"))
		fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
		fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
		fmt.Println(strings.ToLower("Önnek İş"))
		fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

		fmt.Println(strings.ToUpper("hello world"))
		fmt.Println(strings.ToUpper("ā á ǎ à"))
		fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一"))
		fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
		fmt.Println(strings.ToUpper("örnek iş"))
		fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

		fmt.Println(strings.ToTitle("hElLo wOrLd"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
		fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
		fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// string trim
		x := "!!!@@@你好,!@#$ Gophers###$$$"
		fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
		fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
		fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
		fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
		fmt.Println(strings.TrimPrefix(x, "!"))
		fmt.Println(strings.TrimSuffix(x, "$"))

		f := func(r rune) bool {
			return !unicode.Is(unicode.Han, r)
		}
		fmt.Println(strings.TrimFunc(x, f))
		fmt.Println(strings.TrimLeftFunc(x, f))
		fmt.Println(strings.TrimRightFunc(x, f))

		fmt.Println(strings.Repeat("#*", 10))
	}

}
