package main

import (
	"fmt"
	"reflect"
)

type member struct {
	Name string
	Age  int
}

func main() {
	member01 := &member{
		Name: "lily",
		Age:  100,
	}

	fmt.Println("member reflect type :", reflect.TypeOf(member01))          // *main.member
	fmt.Println("*member reflect type :", reflect.TypeOf(*member01))        // main.member
	fmt.Println("*member reflect name :", reflect.TypeOf(*member01).Name()) // member
	fmt.Println("*member reflect kind :", reflect.TypeOf(*member01).Kind()) // struct

	member02 := &member{
		Name: "cyrus",
		Age:  20,
	}

	typeOfMember := reflect.TypeOf(member02).Elem()
	fmt.Println("element name: ", typeOfMember.Name()) // member
	fmt.Println("element kind: ", typeOfMember.Kind()) // struct
}
