package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number int = 1234

	numPoint := reflect.ValueOf(&number)
	num := reflect.ValueOf(number)

	cNumPoint := numPoint.Interface().(*int)
	cNum := num.Interface().(int)

	fmt.Println(cNumPoint, cNum)
}
