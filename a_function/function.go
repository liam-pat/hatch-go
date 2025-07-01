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

func addPrefix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasPrefix(fileName, suffix) {
			return fileName
		}
		return fileName + suffix
	}
}

func manyArgs(args ...int) int {
	return len(args)
}

func main() {
	{
		outputFunc := func() {
			fmt.Printf("name: %s , age : %d\n", "packie", 23)
		}
		outputFunc()

		min, max := func(x int, y int) (min int, max int) {
			if x > y {
				max, min = x, y
			} else {
				max, min = y, x
			}
			return
		}(100, 200)
		fmt.Printf("max: %d ; min: %d ;\n", max, min)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Println("add fun result : ", calc(20, 12, add))
		fmt.Println("sub fun result : ", calc(20, 12, sub))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		function := returnClosure()
		fmt.Println("1st to call>>: ", function())
		fmt.Println("2nd to call>>: ", function())
		fmt.Println("3rd to call>>: ", function())
		fmt.Println("4th to call>>: ", function())
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		for i := 1; i <= 10; i++ {
			fmt.Printf("0 + ...+ %d = %d\n", i, adder()(i))
		}
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Println(addPrefix(".png")("picture"))
		fmt.Println(addPrefix(".jpg")("picture"))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		args := []int{1, 2, 3, 4, 5, 6}
		fmt.Println("args... len: ", manyArgs(args...))
		fmt.Println("args[:2] len: ", manyArgs(args[:2]...))
		fmt.Println("args[2:] len: ", manyArgs(args[2:]...))
		fmt.Println(strings.Repeat("##", 20))
	}
}
