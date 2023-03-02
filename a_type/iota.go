package main

import "fmt"

func main() {

	var (
		name   = "BilYooYam"
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

	// if in same line  they are equal ,diff line it will add one by itself
	const (
		//can guess the type auto
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i:%v ,j1:%v ,j2:%v ,j3:%v , k:%v \n", i, j1, j2, j3, k)

	//若枚举第一个有了公式，下面的元素也会安装这个公式计算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}
