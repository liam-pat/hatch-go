package main

import (
	"fmt"
	"strconv"
)

func main() {

	/**
	ten to bin
	*/
	bin := strconv.FormatInt(13, 2)
	fmt.Println("to bin :", bin)

	/**
	bin to base
	*/
	baseTen, err := strconv.ParseInt("1010", 2, 0)
	if err == nil {
		fmt.Println("to base ten :", baseTen)
	} else {
		fmt.Println("something error :", err)
	}

}
