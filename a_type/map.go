package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	{
		var map01 map[string]string
		map01 = map[string]string{}
		fmt.Println("map01 : ", map01)
		fmt.Println("len: ", len(map01))
		fmt.Println(strings.Repeat("#", 10))
	}
	{

		map02 := make(map[string]string, 2)
		fmt.Println("map02: ", map02)
		fmt.Println("len: ", len(map02))
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		map03 := map[string]string{}
		fmt.Println("map03: ", map03)
		fmt.Println("len: ", len(map03))
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		subjects := make(map[string]string, 2)
		subjects["1st"] = "Java"
		subjects["2nd"] = "C++"
		subjects["3rd"] = "PHP"
		fmt.Println("subjects : ", subjects)

		for key, value := range subjects {
			fmt.Printf("key: %s \t value: %s \n", key, value)
		}

		key := "first"
		if value, has := subjects[key]; has {
			fmt.Printf("key = %s is existing \t value = %s \n", key, value)
		} else {
			fmt.Printf("key = %s is not existing \n", key)
		}
		delete(subjects, key)
		fmt.Println(subjects)
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		var scene sync.Map
		scene.Store("id", 1)
		scene.Store("name", "lily")

		name, _ := scene.Load("name")
		fmt.Println("name: ", name)

		scene.Range(func(k, v interface{}) bool {
			fmt.Println(k, v)
			return true
		})
		scene.Delete("id")
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		newMap := map[string]string{"name": "lily", "age": "30"}
		fmt.Println(newMap)
		changeValue(newMap)
		fmt.Println(newMap)
	}
}

func changeValue(m map[string]string) {
	m["name"] = "james"
}
