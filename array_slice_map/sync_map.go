package main

import (
	"fmt"
	"sync"
)

func main() {

	var scene sync.Map

	//save
	scene.Store("id", 1)
	scene.Store("name", "list")

	//get
	fmt.Println(scene.Load("name"))

	//read
	scene.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
