package main

import (
	"fmt"
	"math"
	"strings"
)

type shaper interface {
	Area() float32
}

type square struct {
	side float32
}

type rectangle struct {
	length, width float32
}

type circle struct {
	radius float32
}

func (sq *square) Area() float32 { return sq.side * sq.side }
func (rg rectangle) Area() float32 {
	return rg.length * rg.width
}
func (ce *circle) Area() float32 {
	return math.Pi * ce.radius * ce.radius
}

func main() {
	{
		r := rectangle{5, 3}
		s := &square{5}
		c := &circle{5}

		shapes := []shaper{r, s, c}
		for _, b := range shapes {
			fmt.Println("data : ", b)
			fmt.Println("area : ", b.Area())
		}
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		var shape shaper
		sq := new(square)
		sq.side = 5
		shape = sq
		if t, ok := shape.(*square); ok {
			fmt.Printf("is %T\n", t)
		}
		if u, ok := shape.(*circle); ok {
			fmt.Printf("is %T\n", u)
		} else {
			fmt.Println("is not circle")
		}
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		var shape shaper
		sq := new(square)
		sq.side = 5
		shape = sq

		switch t := shape.(type) {
		case *square, *circle:
			fmt.Printf("type = %[1]T, value %[1]v\n", t)
		case nil:
			fmt.Printf("nil\n")
		default:
			fmt.Printf("err %T\n", t)
		}
		fmt.Println(strings.Repeat("#", 10))
	}
	{
		var shape shaper
		sq := new(square)
		sq.side = 5
		shape = sq

		if sv, ok := shape.(shaper); ok {
			fmt.Printf("implements interface shaper: %s\n", sv)
		} else {
			fmt.Println("does not implement interface shaper")
		}
		fmt.Println(strings.Repeat("#", 10))
	}

}
