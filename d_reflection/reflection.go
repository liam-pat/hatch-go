package main

import (
	"fmt"
	"reflect"
	"strings"
)

type member struct {
	Name string
	Age  int
}

type student struct {
	Name string
	Sex  string
}

func (obj student) Fn() {
	fmt.Println("this is student 's method")
}

type person struct {
	Name string
	Age  int
}

func (p person) output1() {
	fmt.Println("this is person 's method one")
}

func (p person) output2(param int) {
	fmt.Println("this is person 's method two , parameter", param)
}

func add(name string, age int) {
	fmt.Printf("name is %s, age is %d \n", name, age)
}

func main() {
	{
		var number int = 123
		var interfaceNum interface{} = number

		fmt.Println("interface type: ", reflect.TypeOf(interfaceNum))
		fmt.Println("interface value: ", reflect.ValueOf(interfaceNum))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		var number int64 = 512

		// can change the value of the variable
		obj := reflect.ValueOf(&number)
		fmt.Printf("%-30s: %v\n", "reflect.ValueOf(&number)", obj)
		fmt.Printf("%-30s: %v\n", "obj.Elem()", obj.Elem())
		fmt.Printf("%-30s: %v\n%-30s: %v\n", "obj.Elem().type", obj.Elem().Type(), "can set?", obj.Elem().CanSet())
		obj.Elem().SetInt(1024)
		fmt.Printf("%-30s: %v\n", "obj.Elem().SetInt(1024)", number)

		fmt.Println(strings.Repeat("##", 20))

		// can not change the value of the variable
		obj2 := reflect.ValueOf(number)
		fmt.Printf("%-30s: %v\n", "reflect.ValueOf(number)", obj2)
		fmt.Printf("%-30s: %v\n", "obj2.Int()", obj2.Int())
		fmt.Printf("%-30s: %v\n", "obj2.Interface().(int64)", obj2.Interface().(int64))

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		member01 := &member{Name: "lily", Age: 100}
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(member01)", reflect.TypeOf(member01))                 // *main.member
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(*member01)", reflect.TypeOf(*member01))               // main.member
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(*member01).Name()", reflect.TypeOf(*member01).Name()) // member
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(*member01).Kind()", reflect.TypeOf(*member01).Kind()) // struct

		fmt.Println(strings.Repeat("##", 20))

		mem02 := &member{Name: "cyrus", Age: 20}
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(mem02).Elem().Name()", reflect.TypeOf(mem02).Elem().Name()) // member
		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(mem02).Elem().Kind()", reflect.TypeOf(mem02).Elem().Kind()) // struct

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		st := student{Name: "Joy", Sex: "male"}
		var stInterface interface{} = st

		fmt.Printf("%-40s: %v\n", "reflect.TypeOf(stInterface)", reflect.TypeOf(stInterface))
		fmt.Printf("%-40s: %v\n", "reflect.ValueOf(stInterface)", reflect.ValueOf(stInterface))

		type1 := reflect.TypeOf(stInterface)
		value1 := reflect.ValueOf(stInterface)
		// property
		for i := 0; i < type1.NumField(); i++ {
			field := type1.Field(i)
			value := value1.Field(i)
			fmt.Printf("key: %q, type: %q, value: %q\n", field.Name, field.Type, value)
		}
		// function
		for i := 0; i < type1.NumMethod(); i++ {
			f := type1.Method(i)
			fmt.Printf("method name: %q, type: %q\n", f.Name, f.Type)
		}
		// check the struct property
		if name, ok := type1.FieldByName("Name"); ok {
			fmt.Printf("find the field name : %v\n", name)
		}
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		person := &person{Name: "alan", Age: 25}

		var psinterface interface{} = person
		value := reflect.ValueOf(psinterface)

		method1 := value.MethodByName("output1")
		args01 := make([]reflect.Value, 0)
		method1.Call(args01)

		method2 := value.MethodByName("output2")
		args02 := []reflect.Value{reflect.ValueOf(30)}
		method2.Call(args02)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		// call a function via reflection
		funcValue := reflect.ValueOf(add)
		params := []reflect.Value{reflect.ValueOf("jerry"), reflect.ValueOf(20)}
		reList := funcValue.Call(params)
		fmt.Println(reList)
	}
}
