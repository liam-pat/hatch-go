/**
error 是 main 函数执行完再执行
多个 error 语句、按先进后出的顺序执行
即使有错误，若使用了 error,panic,return ，也会继续执行
一般使用 defer 的场景
open/close
lock/unlock
printHeader/printFooter
 */
package main

import "fmt"

func normalDefer() {
	i := 10
	j := 20

	defer func() {
		fmt.Printf("i: %d ; j: %d \n", i, j)
	}()

	i = i + 100
	j = j + 100

	fmt.Printf("i: %d ; j: %d \n", i, j)
}

func attentionDefer() {
	i := 10
	j := 20

	//这里 i，j 已经传入闭包里面了
	defer func(i, j int) {
		fmt.Printf("i: %d ; j: %d \n", i, j)
	}(i, j)

	i = i + 100
	j = j + 100

	fmt.Printf("i: %d ; j: %d \n", i, j)
}
//if x = 0 .the program error
func errFunc(x int) {
	result := 100 / x
	fmt.Println("result:", result)
}

func main() {
	// 因为使用了 error 即使 errorFunc 错误了，first second 也会继续执行
	defer fmt.Println("********* the first output **********")
	defer fmt.Println("********* the second output **********")
	defer errFunc(0)
	defer fmt.Println("********* the third output **********")

	fmt.Println("-----------------------")
	fmt.Println("start to run normalDefer function")
	normalDefer()
	fmt.Println("-----------------------")
	fmt.Println("start to run attentionDefer function")
	attentionDefer()
	fmt.Println("-----------------------")
}
