package main

import "fmt"

func main() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("Please enter your full name: ")
	// explode ' ' input value to parameter
	fmt.Scanln(&firstName, &lastName)

	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %q %q!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
