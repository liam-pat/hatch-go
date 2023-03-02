package main

import (
	"fmt"
	"strings"
)

type calcFun func(int, int) int

func calc(a, b int, usingFunc calcFun) (result int) {
	return usingFunc(a, b)
}
func add(a, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

func returnClosure() func() int {
	var i int
	return func() int {
		i++
		return i * i
	}
}

func adder() func(int) int {
	sum := 0
	return func(number int) int {
		sum += number
		return sum
	}
}

func makeAddSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasPrefix(fileName, suffix) {
			return fileName
		}
		return fileName + suffix
	}
}

func supportManyArgs(args ...int) int {
	return len(args)
}

func main() {
	{
		// anonymous fun
		age := 23
		name := "BilYooYam"

		outPut := func() {
			fmt.Println("name: ", name)
			fmt.Println("age: ", age)
		}
		outPut()
		min, max := func(x int, y int) (min int, max int) {
			if x > y {
				max = x
				min = y
			} else {
				max = y
				min = x
			}
			return
		}(100, 200)
		fmt.Printf("max: %d ; min: %d ;\n", max, min)
	}
	{
		// callback
		fmt.Println(strings.Repeat("#*", 20))
		fmt.Printf("add fun result : %d \t \nsub fun result : %d \n", calc(20, 12, add), calc(20, 12, sub))
	}
	{
		// closure
		fmt.Println(strings.Repeat("#*", 20))
		func1 := returnClosure()
		fmt.Println("first to use: ", func1())
		fmt.Println("second to use: ", func1())
		fmt.Println("third to use: ", func1())
		fmt.Println("fourth to use: ", func1())
		a := adder()
		for i := 0; i < 10; i++ {
			fmt.Printf("0 + ...+ %d = %d\n", i, a(i))
		}
	}
	{
		// define fun
		fmt.Println(strings.Repeat("#*", 20))
		addPngSuffixFun := makeAddSuffix(".png")
		addJpgSuffixFun := makeAddSuffix(".jpg")
		a := addPngSuffixFun("test")
		b := addJpgSuffixFun("test")

		fmt.Println(a)
		fmt.Println(b)
	}
	{
		// support many args
		fmt.Println(strings.Repeat("#*", 20))
		args := []int{1, 2, 3, 4, 5, 6}
		fmt.Printf("args... : %d \t args[:2] : %d \t args[2:] : %d\n", supportManyArgs(args...), supportManyArgs(args[:2]...), supportManyArgs(args[2:]...))
	}

}
