package main

import "fmt"

func main() {
	age := 23
	name := "BilYooYam"

	outPut := func() {
		fmt.Println("name: ", name)
		fmt.Println("age: ", age)
	}
	outPut()

	min, max := func(x int, y int) (min int, max int) {
		if x > y {
			max = x
			min = y
		} else {
			max = y
			min = x
		}
		return
	}(100, 200)

	fmt.Printf("max: %d ; min: %d ;\n", max, min)
}
