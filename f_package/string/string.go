package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	{
		a, b := "gopher", "hello world"
		fmt.Printf("a = %s, b = %s\n", a, b)
		fmt.Println("strings.Compare(a, b)", strings.Compare(a, b))
		fmt.Println("strings.Compare(a, a)", strings.Compare(a, a))
		fmt.Println("strings.Compare(b, a)", strings.Compare(b, a))

		fmt.Println("strings.EqualFold(a, b)", strings.EqualFold(a, b))
		fmt.Println("strings.EqualFold(a, a)", strings.EqualFold(a, a))
		fmt.Println("strings.EqualFold(b, a)", strings.EqualFold(b, a))

		fmt.Println("strings.EqualFold(GO, go)", strings.EqualFold("GO", "go"))
		fmt.Println("strings.EqualFold(壹, 一)", strings.EqualFold("壹", "一"))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Println("strings.ContainsAny(team, i) = ", strings.ContainsAny("team", "i"))
		fmt.Println("strings.ContainsAny(failure, u & i) =", strings.ContainsAny("failure", "u & i"))
		fmt.Println("strings.ContainsAny(in failure, s g) =", strings.ContainsAny("in failure", "s g"))
		fmt.Println("strings.ContainsAny(foo, ) =", strings.ContainsAny("foo", ""))
		fmt.Println("strings.ContainsAny() =", strings.ContainsAny("", ""))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Printf("strings.Split(\"a,b,c\", \",\") = %q\n", strings.Split("a,b,c", ","))
		fmt.Printf("strings.Split(\"a man a plan a canal panama\", \"a \") = %q\n", strings.Split("a man a plan a canal panama", "a "))
		fmt.Printf("strings.Split(\" xyz \", \"\") = %q\n", strings.Split(" xyz ", ""))
		fmt.Printf("strings.Split(\"\", \"Bernardo O'Higgins\") = %q\n", strings.Split("", "Bernardo O'Higgins"))

		fmt.Println(strings.Repeat("##", 20))
	}

	{
		fmt.Println("strings.HasPrefix(Gopher, Go) =", strings.HasPrefix("Gopher", "Go"))
		fmt.Println("strings.HasPrefix(Gopher, C) =", strings.HasPrefix("Gopher", "C"))
		fmt.Println("strings.HasPrefix(Gopher, ) =", strings.HasPrefix("Gopher", ""))
		fmt.Println("strings.HasSuffix(Amigo, go) =", strings.HasSuffix("Amigo", "go"))
		fmt.Println("strings.HasSuffix(Amigo, Ami) =", strings.HasSuffix("Amigo", "Ami"))
		fmt.Println("strings.HasSuffix(Amigo, ) =", strings.HasSuffix("Amigo", ""))
		fmt.Println(strings.Repeat("##", 20))
	}

	{
		fmt.Println("strings.Replace(\"oink oink oink\", \"k\", \"ky\", 2) =", strings.Replace("oink oink oink", "k", "ky", 2))
		fmt.Println("strings.Replace(\"oink oink oink\", \"oink\", \"moo\", -1) =", strings.Replace("oink oink oink", "oink", "moo", -1))
		fmt.Println("strings.ReplaceAll(\"oink oink oink\", \"oink\", \"moo\") =", strings.ReplaceAll("oink oink oink", "oink", "moo"))
		fmt.Println(strings.Repeat("##", 20))
	}

	{
		fmt.Println("strings.ToLower(\"HELLO WORLD\") =", strings.ToLower("HELLO WORLD"))
		fmt.Println("strings.ToLower(\"Ā Á Ǎ À\") =", strings.ToLower("Ā Á Ǎ À"))
		// support the specific case, eg. turkish
		fmt.Println("strings.ToLowerSpecial(unicode.TurkishCase, \"壹\") =", strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
		fmt.Println("strings.ToLowerSpecial(unicode.TurkishCase, \"HELLO WORLD\") =", strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
		fmt.Println("strings.ToLower(\"Önnek İş\") =", strings.ToLower("Önnek İş"))
		fmt.Println("strings.ToLowerSpecial(unicode.TurkishCase, \"Önnek İş\") =", strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

		fmt.Println("strings.ToUpper(\"hello world\") =", strings.ToUpper("hello world"))
		fmt.Println("strings.ToUpper(\"ā á ǎ à\") =", strings.ToUpper("ā á ǎ à"))

		fmt.Println("strings.ToUpperSpecial(unicode.TurkishCase, \"一\") =", strings.ToUpperSpecial(unicode.TurkishCase, "一"))
		fmt.Println("strings.ToUpperSpecial(unicode.TurkishCase, \"hello world\") =", strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
		fmt.Println("strings.ToUpper(\"örnek iş\") =", strings.ToUpper("örnek iş"))
		fmt.Println("strings.ToUpperSpecial(unicode.TurkishCase, \"örnek iş\") =", strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
		fmt.Println(strings.Repeat("##", 20))

	}

	{
		fmt.Printf("%-70s = %s\n", "strings.Title(\"hElLo wOrLd\")", strings.Title("hElLo wOrLd"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitle(\"hElLo wOrLd\")", strings.ToTitle("hElLo wOrLd"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitleSpecial(unicode.TurkishCase, \"hElLo wOrLd\")", strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))

		fmt.Printf("%-70s = %s\n", "strings.Title(\"āáǎà ōóǒò êēéěè\")", strings.Title("āáǎà ōóǒò êēéěè"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitle(\"āáǎà ōóǒò êēéěè\")", strings.ToTitle("āáǎà ōóǒò êēéěè"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitleSpecial(unicode.TurkishCase, \"āáǎà ōóǒò êēéěè\")", strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))

		fmt.Printf("%-70s = %s\n", "strings.Title(\"dünyanın ilk borsa yapısı Aizonai kabul edilir\")", strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitle(\"dünyanın ilk borsa yapısı Aizonai kabul edilir\")", strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Printf("%-70s = %s\n", "strings.ToTitleSpecial(unicode.TurkishCase, \"dünyanın ilk borsa yapısı Aizonai kabul edilir\")", strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
		fmt.Println(strings.Repeat("##", 20))
	}
}
