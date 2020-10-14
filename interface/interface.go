package main

import "fmt"

type Human interface {
	sayHello()
}

type OldMan struct {
	number int
	name   string
	age    int
}

type younger struct {
	number int
	name   string
	age    int
}

type SomeOne string

func (tmp OldMan) sayHello() {
	fmt.Printf("I am a old man, name = %s 、 age = %d 、name = %s \n", tmp.name, tmp.age, tmp.name)
}

func (tmp younger) sayHello() {
	fmt.Printf("I am a younger, name = %s 、 age = %d 、name = %s \n", tmp.name, tmp.age, tmp.name)
}

func (tmp SomeOne) sayHello() {
	fmt.Printf("I just a stranger\n")
}

func whoSayHello(tmp Human)  {
	tmp.sayHello()
}

func main() {
	var human Human

	oldManObjAddress := &OldMan{1, "old man", 67}
	human = oldManObjAddress
	human.sayHello()

	youngerObj := younger{2, "younger", 21}
	human = &youngerObj
	human.sayHello()

	strangerAddress := SomeOne("strange")
	human = &strangerAddress
	human.sayHello()

	testOldMan := OldMan{3,"test old man",88}
	whoSayHello(testOldMan)
}
