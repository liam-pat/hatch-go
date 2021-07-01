package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func (p person) Fn() {
	fmt.Println("this is person 's method")
}

func main() {
	person01 := &person{
		Name: "human",
		Age:  25,
	}

	var interfacePerson interface{} = person01

	reflectValue := reflect.ValueOf(interfacePerson)

	method1 := reflectValue.MethodByName("Fn")

	// not need arg
	args := make([]reflect.Value, 0)
	method1.Call(args)
}
