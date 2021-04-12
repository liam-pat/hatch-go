package main

import "fmt"

type Simpler interface {
	Get() string
	Set(name string)
}

type person struct {
	name   string
	number int
}

func (obj *person) Get() string {
	return obj.name
}

func (obj *person) Set(name string) {
	obj.name = name
}

type rSimple struct {
	nickname string
}

func (rSimple *rSimple) Get() string {
	return rSimple.nickname
}

func (rSimple *rSimple) Set(nickname string) {

	rSimple.nickname = nickname
}

func main() {
	person1 := &person{"packie", 111}
	person1.Set("packie-copy")
	fmt.Printf("%v \n", person1.Get())
	fmt.Println("###################")

	// assert the interface
	var simplerInter Simpler
	per1 := new(person)
	per1.name = "biyongyao"
	simplerInter = per1
	if assertType, ok := simplerInter.(*person); ok {
		fmt.Printf("The type of simplerInter is: %T\n", assertType)
	}

	// switch to assert type
	var rSimpleInter Simpler
	rSimpleExample := new(rSimple)
	rSimpleExample.nickname = "biyongyao"
	rSimpleInter = rSimpleExample

	switch rSimpleInter.(type) {
	case *person:
		fmt.Println("is *person")
	case *rSimple:
		fmt.Println("is *rSimple")
	default:
		fmt.Println("nothing")
	}
}
