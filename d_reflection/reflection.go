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

func (p person) Fn01() {
	fmt.Println("this is person 's method one")
}

func (p person) Fn02(param int) {
	fmt.Println("this is person 's method two , parameter", param)
}

func add(name string, age int) {
	fmt.Printf("name is %s, age is %d \n", name, age)
}

func main() {
	{
		var number int = 123
		var interfaceNum interface{} = number

		fmt.Println("type: ", reflect.TypeOf(interfaceNum))
		fmt.Println("value: ", reflect.ValueOf(interfaceNum))
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		var number int64 = 233

		reflectObj := reflect.ValueOf(&number)
		value := reflectObj.Elem()

		fmt.Printf("reflect.ValueOf(&number) : %v\n", reflectObj)
		fmt.Printf("reflectObj.Elem() : %v\n", value)
		fmt.Printf("type : %v\t can set? %v\t\n", value.Type(), value.CanSet())

		value.SetInt(111)
		fmt.Printf("value.SetInt(111) and number : %v\n", number)

		reflectObj2 := reflect.ValueOf(number)
		fmt.Println("reflectObj2 get number value", reflectObj2.Int())
		fmt.Println("reflectObj2 get number value", reflectObj2.Interface().(int64))

		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		member01 := &member{Name: "lily", Age: 100}

		fmt.Println("member reflect type :", reflect.TypeOf(member01))          // *main.member
		fmt.Println("*member reflect type :", reflect.TypeOf(*member01))        // main.member
		fmt.Println("*member reflect name :", reflect.TypeOf(*member01).Name()) // member
		fmt.Println("*member reflect kind :", reflect.TypeOf(*member01).Kind()) // struct

		member02 := &member{Name: "cyrus", Age: 20}

		typeOfMember := reflect.TypeOf(member02).Elem()
		fmt.Println("element name: ", typeOfMember.Name()) // member
		fmt.Println("element kind: ", typeOfMember.Kind()) // struct

		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		st := student{Name: "Joy", Sex: "male"}
		var studentInterface interface{} = st

		typeOf := reflect.TypeOf(studentInterface)
		fmt.Println("type of ", typeOf)

		valueOf := reflect.ValueOf(studentInterface)
		fmt.Println("value of ", valueOf)

		// property
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			value := valueOf.Field(i).Interface()
			fmt.Printf("-- field's name: %q, type: %q, value: %q\n", field.Name, field.Type, value)
		}
		// function
		for i := 0; i < typeOf.NumMethod(); i++ {
			m := typeOf.Method(i)
			fmt.Printf("-- method'name: %q, type: %q\n", m.Name, m.Type)
		}
		// check the struct property
		if userName, ok := typeOf.FieldByName("Name"); ok {
			fmt.Printf("existing field name : %v\n", userName)
		}
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		person := &person{Name: "alan", Age: 25}
		var interfacePerson interface{} = person

		valueOf := reflect.ValueOf(interfacePerson)

		// get the Fn01 function
		method1 := valueOf.MethodByName("Fn01")
		args01 := make([]reflect.Value, 0)
		method1.Call(args01)

		// get the Fn02 function
		method2 := valueOf.MethodByName("Fn02")
		args02 := []reflect.Value{reflect.ValueOf(30)}
		method2.Call(args02)

		funcValue := reflect.ValueOf(add)
		params := []reflect.Value{reflect.ValueOf("jerry"), reflect.ValueOf(20)}
		reList := funcValue.Call(params)
		fmt.Println(reList)
	}
}
