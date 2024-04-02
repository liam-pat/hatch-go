package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// will clone new one
func modify1(p [5]int) {
	p[0] = 6
	fmt.Println("modify p =", p)
}

// will pass the &p
func modify2(p *[5]int) {
	(*p)[0] = 6
	fmt.Println("modify *p = ", *p)
}

func main() {
	{
		b := [5]int{1, 2, 3, 4, 5}
		fmt.Println("b = ", b)

		c := [5]int{1, 2, 3}
		fmt.Println("c = ", c)

		d := [5]int{1: 2, 3: 5}
		fmt.Println("d = ", d)

		e := [3][4]int{{1, 2, 3}, {2, 4, 6}}
		fmt.Println("e = ", e)

		a := [5]int{1, 2, 3, 4, 5}
		modify1(a)
		fmt.Println("a = ", a)

		modify2(&a)
		fmt.Println("&a = ", a)
		/**
		  output :
		  b =  [1 2 3 4 5]
		  c =  [1 2 3 0 0]
		  d =  [0 2 0 5 0]
		  e =  [[1 2 3 0] [2 4 6 0] [0 0 0 0]]
		  modify p = [6 2 3 4 5]
		  a =  [1 2 3 4 5]
		  modify *p =  [6 2 3 4 5]
		  &a =  [6 2 3 4 5]
		*/
	}
	{
		/**
		array slice map type data will create new one
		*/
		nums := [3]int8{11, 22, 33}
		cloneOne := nums

		fmt.Printf("nums address : %p \n", &nums)
		fmt.Printf("ponit to anthor one address : %p \n", &cloneOne)

		fmt.Printf("array address : %p \n", &nums)
		fmt.Printf("array[1] address : %p \n", &nums[0])
		fmt.Printf("array[2] address : %p \n", &nums[1])
		fmt.Printf("array[3] address : %p \n", &nums[2])
		/**
		output :
		nums address : 0x14000122002
		demo address : 0x14000122005
		array_slice_map address : 0x14000122002
		array_slice_map[1] address : 0x14000122002
		array_slice_map[2] address : 0x14000122003
		array_slice_map[3] address : 0x14000122004
		*/
	}
	{
		fmt.Println("#################")
		var books [4]string

		books[0] = "Kafka's Revenge"
		books[1] = "Stay Golden"
		books[2] = "Everything"
		books[3] += books[0] + " 2nd Edition"
		//%#v include the array type
		fmt.Printf("books  : %#v\n", books)
	}
	{
		fmt.Println("#################")
		books := [...]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everything",
			"Kafka's Revenge 2nd Edition",
		}
		fmt.Printf("books : %#v\n", books)
	}
	// output the emoji in random
	{
		fmt.Println("#################")
		args := os.Args[1:]

		var name string
		if len(args) != 1 {
			name = "test"
		} else {
			name = "test"
		}
		moods := [...]string{"happy 😀", "good 👍", "awesome 😎", "sad 😞", "bad 👎", "terrible 😩"}

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods))

		fmt.Printf("%s feels %s\n", name, moods[n])
	}
	{
		type placeholder [5]string
		zero := placeholder{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		}

		one := placeholder{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		}
		digits := [...]placeholder{zero, one}

		for line := range digits[0] {
			for digit := range digits {
				fmt.Print(digits[digit][line], "  ")
			}
			fmt.Println()
		}
	}

}
