package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number int64 = 233

	// pass pointer
	reflectValue := reflect.ValueOf(&number)
	newValue := reflectValue.Elem()
	fmt.Println("type of peElem: ", newValue.Type())
	fmt.Println("can set?", newValue.CanSet())
	newValue.SetInt(111)
	fmt.Println("after change", number)

	// get value
	reflectValue = reflect.ValueOf(number)
	fmt.Println("reflect get number value", reflectValue.Int())
	fmt.Println("reflect get number value", reflectValue.Interface().(int64))
}
