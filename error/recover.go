package main

//1. 仅在 defer 调用
//2. 获取 panic 的值
//3. 如果没法处理，可重新 panic

import "fmt"

func test01() {
	fmt.Println("test A")
}

func test02(index int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [5]int
	arr[index] = 100
	fmt.Println("test B")
}

func test03() {
	fmt.Println("test C")
}

func main() {
	test01()
	test02(10)
	test03()
}
