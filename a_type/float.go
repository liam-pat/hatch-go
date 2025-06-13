package main

import "fmt"

func main() {

	{
		fmt.Printf("result: %[1]T  \t %[1]v\n", 3.1415926)
		fmt.Printf("result: %10.2f \n", 3.1415926)
		fmt.Printf("result: %010.2f\n", 3.1415926)
	}
}
