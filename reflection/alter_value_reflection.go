package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number int = 233

	value := reflect.ValueOf(&number)
	pElem := value.Elem()

	fmt.Println("type of peElem: ", pElem.Type())
	fmt.Println("can set?", pElem.CanSet())

	pElem.SetInt(111)
	fmt.Println("after change", number)
}
