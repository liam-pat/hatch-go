package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number int = 123

	var interfaceNum interface{} = number

	fmt.Println("type: ", reflect.TypeOf(interfaceNum))
	fmt.Println("value: ", reflect.ValueOf(interfaceNum))
}
