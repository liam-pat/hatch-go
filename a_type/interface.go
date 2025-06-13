package main

import "fmt"

type Human interface {
	sayHello()
}

type older struct {
	number int
	name   string
	age    int
}

type younger struct {
	number int
	name   string
	age    int
}

type someOne string

func (tmp older) sayHello() {
	fmt.Printf("older, name = %s | age = %d \n", tmp.name, tmp.age)
}

func (tmp younger) sayHello() {
	fmt.Printf("younger, name = %s | age = %d \n", tmp.name, tmp.age)
}

func (tmp someOne) sayHello() {
	fmt.Printf("stranger, name = %s \n", tmp)
}

func main() {
	var human Human

	o := &older{1, "older", 67}
	human = o
	human.sayHello()

	y := younger{2, "younger", 21}
	human = &y
	human.sayHello()

	strangerAddress := someOne("strange")
	human = &strangerAddress
	human.sayHello()
}
