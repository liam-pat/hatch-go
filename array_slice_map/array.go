package main

import "fmt"

func modify1(p [5]int) {
	p[0] = 6
	fmt.Println("through modify1 fun p =", p)
}

func modify2(p *[5]int) {
	(*p)[0] = 6
	fmt.Println("through modify2 fun *p = ", *p)
}

func main() {
	//1: init all
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b= ", b)

	//2: init part
	c := [5]int{1, 2, 3}
	fmt.Println("c= ", c)

	//3.init what you want
	d := [5]int{1: 2, 3: 5}
	fmt.Println("d= ", d)

	//if you do not init it ,default 0
	e := [3][4]int{{1, 2, 3}, {2, 4, 6}}
	fmt.Println("e= ", e)
	fmt.Println("------------------------------------------------------")

	a := [5]int{1, 2, 3, 4, 5}
	modify1(a)
	fmt.Println("outside modify1 fun, main a  = ", a)
	fmt.Println("------------------------------------------------------")
	modify2(&a)
	fmt.Println("outside modify2 fun, main a  = ", a)
}
