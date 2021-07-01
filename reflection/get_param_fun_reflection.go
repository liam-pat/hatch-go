package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Sex  string
}

func (obj student) Fn() {
	fmt.Println("this is student 's method")
}

func main() {
	st := &student{
		Name: "Joy",
		Sex:  "male",
	}

	var studentInterface interface{} = st

	interfaceType := reflect.TypeOf(studentInterface)
	fmt.Println("studentInterface 's type", interfaceType.Name())

	interfaceValue := reflect.ValueOf(studentInterface)
	fmt.Println("studentInterface 's value", interfaceValue)

	// get the parameters
	for i := 0; i < interfaceType.NumField(); i++ {
		field := interfaceType.Field(i)
		value := interfaceValue.Field(i).Interface()
		fmt.Printf("field's name: %s, type: %v, value: %v\n", field.Name, field.Type, value)
	}

	// get functions
	for i := 0; i < interfaceType.NumMethod(); i++ {
		m := interfaceType.Method(i)
		fmt.Printf("method'name: %s, type: %v\n", m.Name, m.Type)
	}

}
