package main

import "fmt"

func getParamsLength(args ...int) int {
	fmt.Printf("have %d param \n", len(args))
	return len(args)
}

func getParamsLength1(args ...int) (length int) {
	length = len(args)
	return
}

func returnMore() (name string, age int, gender byte) {
	name = "BilYooYam"
	age = 23
	gender = 'M'
	return
}

func getSum(number int) int {
	if number == 1 {
		return 1
	}

	return getSum(number-1) + number
}

func main() {
	getParamsLength(1, 2, 4, 5, 6)
	fmt.Println("length:", getParamsLength1(1, 2, 4, 5, 6))

	name, age, gender := returnMore()
	fmt.Printf("name: %s ; age: %d ; gender: %c \n", name, age, gender)

	fmt.Println("1+2+3+4+...+100 =",getSum(100))
}
