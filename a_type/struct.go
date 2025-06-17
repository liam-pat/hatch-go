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
		fmt.Println("student01: ", student01)

		student02 := Student{Name: "girlFriend", Age: 27}
		fmt.Println("student02: ", student02)

		//struct point init
		var student03 Student
		var studentPoint *Student
		studentPoint = &student03
		studentPoint.Age = 18
		studentPoint.Name = "poiter"
		fmt.Printf("student03: %v , studentPoint: %v\n", student03, studentPoint)

		var student04 = new(Student)
		student04.Name = "new_student"
		student04.Age = 50
		fmt.Println("student04: ", student04)
		fmt.Println(strings.Repeat("***", 20))
	}
	{
		users := []Student{
			{1, "name1", "man", 21, "GZ"},
			{2, "name2", "girl", 22, "SZ"},
		}
		out, _ := json.MarshalIndent(users, "<<<", "  ")
		fmt.Printf("output type: %T,json encode: %s\n", out, string(out))

		fmt.Println(strings.Repeat("***", 20))
	}
	{
		// json decode
		type user struct {
			Name        string          `json:"username"`
			Permissions map[string]bool `json:"perms"`
		}
		jsonStr := `[
	{"username": "lily"},
	{"username": "god","perms": {"admin": true}},
	{"username": "devil","perms": {"write": true}}
]`
		var users []user
		if err := json.Unmarshal([]byte(jsonStr), &users); err != nil {
			fmt.Println(err)
			return
		}

		for _, user := range users {
			switch p := user.Permissions; {
			case p == nil:
				fmt.Printf("%s no permission\n", user.Name)
			case p["admin"]:
				fmt.Printf("%s is an admin\n", user.Name)
			case p["write"]:
				fmt.Printf("%s is a write\n", user.Name)
			}
		}
		fmt.Println(strings.Repeat("***", 20))
	}
}
