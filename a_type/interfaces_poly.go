package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (rg Rectangle) Area() float32 {
	return rg.length * rg.width
}

func (ce *Circle) Area() float32 {
	return math.Pi * ce.radius * ce.radius
}

func main() {
	r := Rectangle{5, 3}
	s := &Square{5}
	c := &Circle{5}

	shapes := []Shaper{r, s, c}

	for _, b := range shapes {
		fmt.Println("shape data : ", b)
		fmt.Println("shape area : ", b.Area())
		fmt.Println("-----------")
	}

	// assert interface 类型断言
	var areaInterface Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaInterface = sq1
	if t, ok := areaInterface.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaInterface.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaInterface does not contain a variable of type Circle")
	}
	fmt.Println("######")

	// 类型判断
	switch t := areaInterface.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	fmt.Println("######")

	/*
		# 判断变量是否实现了某个接口

		if sv, ok := v.(Stringer); ok {
			fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
		}
	*/
}
