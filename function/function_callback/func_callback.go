package main

import "fmt"

//same as the interface
type calcTypeFunc func(int, int) int

func Calc(a, b int, useFunc calcTypeFunc) (result int) {
	fmt.Println("start to calc");
	result = useFunc(a, b)
	return
}

func add(a, b int) int {
	fmt.Println("you use the add func")
	return a + b
}

func main() {
	Calc(10,12,add)
}
