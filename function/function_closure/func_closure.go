package main

import "fmt"

//表示返回一个函数，这个函数里面再返回一个整数
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

func main() {
	//返回一个闭包函数
	//因为闭包还在使用 i 这个变量，所以这个一直存在
	func1 := returnClosure()

	fmt.Println("first to use: ", func1())
	fmt.Println("second to use: ", func1())
	fmt.Println("third to use: ", func1())
	fmt.Println("fourth to use: ", func1())

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ...+ %d = %d\n",i,a(i))
	}
}
