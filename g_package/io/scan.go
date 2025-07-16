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
	fmt.Println("plz input first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("plz input last name: ")
	fmt.Scanln(&lastName)

	fmt.Printf("Hi %q %q!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("read: ", f, i, s)
}
