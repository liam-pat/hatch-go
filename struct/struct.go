package main

import "fmt"

type Student struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

func main() {
	//init the struct
	var student01 Student = Student{1, "author", "male", 23, "GuangZhou",}
	//	var student01 *Student = &Student{1, "author", "male", 23, "GuangZhou",}
	fmt.Println("--------------------------------------------")
	fmt.Println("student01 ", student01)
	fmt.Println("--------------------------------------------")

	student02 := Student{name: "girlFriend", age: 27}
	//	student02 := &Student{name: "girlFriend", age: 27}
	fmt.Println("student02 ", student02)
	fmt.Println("--------------------------------------------")

	//struct point init
	var student03 Student
	var studentPoint *Student
	studentPoint = &student03
	studentPoint.age = 18
	studentPoint.name = "test"
	fmt.Println("studentPoint ", studentPoint)
	fmt.Println("--------------------------------------------")

	var student04 = new(Student)
	student04.name = "initByNew"
	student04.age = 50
	fmt.Println("student04", student04)
	fmt.Println("--------------------------------------------")
}
