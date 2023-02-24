package main

import (
	"fmt"
	"sync"
)

// map 打印出来可能是乱序的 ,map 只有 len 没有 cap，

func main() {
	{
		// case 1) : define new one
		var map01 map[string]string
		fmt.Println("map01 : ", map01)
		fmt.Printf("map01 len : %d \n", len(map01))

		// case 2) : define new one
		map02 := make(map[string]string, 2)
		map02["first"] = "Java"
		map02["second"] = "C++"
		map02["third"] = "PHP"
		fmt.Println("map02 : ", map02)
		fmt.Printf("map02 len : %d \n", len(map02))

		// get the all elements
		for key, value := range map02 {
			fmt.Printf("map02 key: %s \t value: %s \n", key, value)
		}

		// get the element
		key := "first"
		if value, isExist := map02[key]; isExist {
			fmt.Printf("key = %s is existing \t value = %s \n", key, value)
		} else {
			fmt.Printf("key = %s is not existing \n", key)
		}

		// delete the element
		delete(map02, key)
		fmt.Println(map02)
	}
	/**
	output :
	map01 :  map[]
	map01 len : 0
	map02 :  map[first:Java second:C++ third:PHP]
	map02 len : 3
	map02 key: first 	 value: Java
	map02 key: second 	 value: C++
	map02 key: third 	 value: PHP
	key = first is existing 	 value = Java
	map[second:C++ third:PHP]
	*/
	{
		var scene sync.Map
		scene.Store("id", 1)
		scene.Store("name", "lily")

		fmt.Println(scene.Load("name"))

		scene.Range(func(k, v interface{}) bool {
			fmt.Println(k, v)
			return true
		})
	}

	{
		newMap := make(map[string]string, 2)
		newMap["name"] = "lily"
		newMap["age"] = "30"
		fmt.Println(newMap)
		TestMapFun(newMap)
		fmt.Println(newMap)
		/**
		output :
		map[age:30 name:lily]
		map[age:30 name:james]
		map[age:30 name:james]
		*/
	}
}

func TestMapFun(m map[string]string) {
	m["name"] = "james"
	fmt.Println(m)
}
