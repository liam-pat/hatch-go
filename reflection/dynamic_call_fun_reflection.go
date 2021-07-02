package main

import (
	"fmt"
	"reflect"
)

// common function
func add(name string, age int) {
	fmt.Printf("name is %s, age is %d \n", name, age)
}

type person struct {
	Name string
	Age  int
}

func (p person) Fn01() {
	fmt.Println("this is person 's method one")
}

func (p person) Fn02(param int) {
	fmt.Println("this is person 's method two , parameter", param)
}
func main() {
	person01 := &person{Name: "human", Age: 25}
	var interfacePerson interface{} = person01
	reflectValue := reflect.ValueOf(interfacePerson)

	method1 := reflectValue.MethodByName("Fn01")
	args01 := make([]reflect.Value, 0) // empty slice as same as do not has args
	method1.Call(args01)

	method2 := reflectValue.MethodByName("Fn02")
	args02 := []reflect.Value{reflect.ValueOf(30)}
	method2.Call(args02)

	// common function reflect call
	funcValue := reflect.ValueOf(add)
	params := []reflect.Value{reflect.ValueOf("jerry"), reflect.ValueOf(20)}
	reList := funcValue.Call(params)
	fmt.Println(reList)
}
