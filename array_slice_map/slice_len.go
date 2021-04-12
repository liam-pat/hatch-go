package main

import "fmt"

func main() {

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]

	b := a[1:1]
	c := a[1:2]

	fmt.Println("b length  ", len(b))
	fmt.Println("c length  ", len(c))

}
