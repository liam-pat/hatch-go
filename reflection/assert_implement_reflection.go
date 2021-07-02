package main

import (
	"fmt"
	"reflect"
)

type tester interface {
	test()
}

type qa struct {
	Name string
}

func (tester *qa) test() {
	fmt.Printf("I am a %s,testing", tester.Name)
}

func main() {
	qaStruct := &qa{Name: "qa-member-01"}

	testerInterface := reflect.TypeOf((*tester)(nil)).Elem()
	reflectQaStruct := reflect.TypeOf(qaStruct)

	fmt.Println(reflectQaStruct.Implements(testerInterface))
}
