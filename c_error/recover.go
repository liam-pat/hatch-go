package main

//1. 仅在 defer 调用
//2. 获取 panic 的值
//3. 如果没法处理，可重新 panic

import (
	"fmt"
	"strings"
)

func test01() {
	fmt.Println("test fun A")
}

func test02(index int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("print err: ", err)
		}
	}()
	var arr [5]int
	arr[index] = 100
	fmt.Println("test fun B")
}

func test03() {
	fmt.Println("test fun C")
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("err msg:", err)
		} else {
			fmt.Println("unknown error")
		}
	}()
	panic("custom error")
}

func main() {
	{
		test01()
		test02(10)
		test03()

		fmt.Println(strings.Repeat("##", 20))

	}
	{
		tryRecover()
	}
}
