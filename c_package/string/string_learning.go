package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	{
		a := "gopher"
		b := "hello world"
		fmt.Println(strings.Compare(a, b))
		fmt.Println(strings.Compare(a, a))
		fmt.Println(strings.Compare(b, a))

		fmt.Println(strings.EqualFold("GO", "go"))
		fmt.Println(strings.EqualFold("壹", "一"))
		fmt.Println(strings.Repeat("#", 20))
	}
	{
		fmt.Println(strings.ContainsAny("team", "i"))
		fmt.Println(strings.ContainsAny("failure", "u & i"))
		fmt.Println(strings.ContainsAny("in failure", "s g"))
		fmt.Println(strings.ContainsAny("foo", ""))
		fmt.Println(strings.ContainsAny("", ""))
		fmt.Println(strings.Repeat("#", 20))
	}
	{
		fmt.Printf("%q\n", strings.Split("a,b,c", ","))
		fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
		fmt.Printf("%q\n", strings.Split(" xyz ", ""))
		fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

		fmt.Println(strings.Repeat("#", 20))
	}

	{
		fmt.Println(strings.HasPrefix("Gopher", "Go"))
		fmt.Println(strings.HasPrefix("Gopher", "C"))
		fmt.Println(strings.HasPrefix("Gopher", ""))
		fmt.Println(strings.HasSuffix("Amigo", "go"))
		fmt.Println(strings.HasSuffix("Amigo", "Ami"))
		fmt.Println(strings.HasSuffix("Amigo", ""))
		fmt.Println(strings.Repeat("#", 20))
	}

	{
		fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
		fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
		fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
		fmt.Println(strings.Repeat("#", 20))
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
		fmt.Println(strings.Repeat("#", 20))

	}

	{
		fmt.Println(strings.Title("hElLo wOrLd"))
		fmt.Println(strings.ToTitle("hElLo wOrLd"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
		fmt.Println(strings.Title("āáǎà ōóǒò êēéěè"))
		fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
		fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.Repeat("#", 20))
	}
}
