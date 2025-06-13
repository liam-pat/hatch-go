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
type rSimple struct {
	nickname string
}

func (obj *person) Get() string {
	return obj.name
}

func (obj *person) Set(name string) {
	obj.name = name
}

func (rSimple *rSimple) Set(nickname string) {
	rSimple.nickname = nickname
}

func (rSimple *rSimple) Get() string {
	return rSimple.nickname
}

func main() {
	{
		p := &person{"packie", 111}
		p.Set("packie-copy")
		fmt.Printf("%v \n", p.Get())
	}
	{
		var simplerInter Simpler
		per1 := new(person)
		per1.name = "liam"
		simplerInter = per1

		if _, ok := simplerInter.(*person); ok {
			fmt.Println("is *person")
		} else {
			fmt.Println("is not *person")
		}
	}
	{
		var rSimpleInter Simpler

		rSimpleExample := new(rSimple)
		rSimpleExample.nickname = "packie"

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
}
