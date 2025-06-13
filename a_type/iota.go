package main

import "fmt"

func main() {

	{
		var (
			name   = "liam"
			age    = 23
			gender = 'M'
		)
		fmt.Printf("name: %v ; age: %d ;gender: %c \n", name, age, gender)
		const (
			cpp = iota
			java
			php
			goland
			javascript
		)
		fmt.Println(cpp, java, php, goland, javascript)

		const (
			i          = iota
			j1, j2, j3 = iota, iota, iota
			k          = iota
		)
		fmt.Printf("i:%v ,j1:%v ,j2:%v ,j3:%v , k:%v \n", i, j1, j2, j3, k)

		const (
			b = 1 << (10 * iota)
			kb
			mb
			gb
			tb
		)
		fmt.Println(b, kb, mb, gb, tb)
	}
	{
		const (
			monday = iota
			tuesday
			wednesday
			thursday
			friday
			saturday
			sunday
		)
		fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
		const (
			monday01 = iota + 1
			tuesday01
			wednesday01
			thursday01
			friday01
			saturday01
			sunday01
		)
		fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
		const (
			EST = -(5 + iota) // CORRECT: -5
			_
			MST // CORRECT: -7
			PST // CORRECT: -8
		)
		fmt.Println(EST, MST, PST)

		const (
			_   = iota // 0
			Jan        // 1
			Feb        // 2
			Mar        // 3
		)
		fmt.Println(Jan, Feb, Mar)

		const (
			Spring = (iota + 1) * 3
			Summer
			Fall
			Winter
		)
		fmt.Println(Spring, Summer, Fall, Winter)
	}
}
