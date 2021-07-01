package main

import "fmt"

func test(s []int) {
	fmt.Printf("test---%p\n", s) // as same as main
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Printf("test---%p\n", s) // over , it will add the cap auto
	fmt.Println("test---", s)    // [0 0 0 1 2 3 4 5]
}

func main() {
	s1 := make([]int, 3)
	test(s1)
	fmt.Printf("main---%p\n", s1) // still the s1 , cap and len also equal 3
	fmt.Println("main---", s1)    // [ 0 0 0]
}
