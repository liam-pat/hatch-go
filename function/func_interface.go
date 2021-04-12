package main

import "fmt"

/**
go语言巧妙使用 Interface
*/
type calcFunc func(int, int) int

func Add(number1 int, number2 int) int {
	println("method add")
	return number1 + number2
}

func Subtract(number1 int, number2 int) int {
	println("method subtract")
	return number1 - number2
}

func calc(a, b int, usingFun calcFunc) (result int) {
	fmt.Println("--------calc interface using--------")
	result = usingFun(a, b)
	fmt.Println("--------calc interface ending--------")
	fmt.Printf("\n")
	return
}

func main() {
	addResult := calc(1, 1, Add)
	subtractResult := calc(100, 99, Subtract)

	fmt.Printf("add result is %d\nsubtract result is %d\n", addResult, subtractResult)
}
