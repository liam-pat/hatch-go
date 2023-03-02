package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	{
		var student01 Student = Student{1, "author", "male", 23, "GZ"}
		fmt.Println(strings.Repeat("***", 10))
		fmt.Println("student01 : ", student01)
		fmt.Println(strings.Repeat("***", 10))

		student02 := Student{Name: "girlFriend", Age: 27}
		fmt.Println("student02 : ", student02)
		fmt.Println(strings.Repeat("***", 10))

		//struct point init
		var student03 Student
		var studentPoint *Student
		studentPoint = &student03
		studentPoint.Age = 18
		studentPoint.Name = "test_point"
		fmt.Println("studentPoint : ", studentPoint)
		fmt.Println(strings.Repeat("***", 10))

		var student04 = new(Student)
		student04.Name = "initByNew"
		student04.Age = 50
		fmt.Println("student04 : ", student04)
		fmt.Println(strings.Repeat("***", 10))
	}
	{
		// json encode . the property need public
		users := []Student{
			{1, "name1", "man", 21, "GZ"},
			{2, "name2", "girl", 22, "SZ"},
			{3, "name3", "nil", 23, "BJ"},
		}
		out, err := json.MarshalIndent(users, "", "\t")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(out))
	}
	{
		// json decode
		type user struct {
			Name        string          `json:"username"`
			Permissions map[string]bool `json:"perms"`
		}
		jsonStr := `[
	{"username": "inanc"},
	{"username": "god","perms": {"admin": true}},
	{"username": "devil","perms": {"write": true}}
]`
		var users []user
		if err := json.Unmarshal([]byte(jsonStr), &users); err != nil {
			fmt.Println(err)
			return
		}

		for _, user := range users {
			fmt.Print("+++", user.Name)

			switch p := user.Permissions; {
			case p == nil:
				fmt.Print(" has no power.\n")
			case p["admin"]:
				fmt.Print(" is an admin.\n")
			case p["write"]:
				fmt.Print(" can write.\n")
			}
		}

	}
}
