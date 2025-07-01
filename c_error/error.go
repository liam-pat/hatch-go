package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func div(a, b float32) (result float32, err error) {
	err = nil
	if b == 0 {
		result = 0
		err = errors.New("Divisor can not be 0")
	} else {
		result = a / b
	}
	return
}
func main() {
	{
		result1, err1 := div(10, 20)
		if err1 != nil {
			fmt.Println(err1)
		} else {
			fmt.Printf("result = %f \n", result1)
		}

		result2, err2 := div(10, 0)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Printf("result = %f \n", result2)
		}
	}
}
