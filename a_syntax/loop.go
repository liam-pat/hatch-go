package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {

	{
		str := "foreach one"
		i := 1

		number := 10000000
		better := 10_000_000

		fmt.Println(number, "-", better)

		for key, value := range str {
			fmt.Printf("key: %d ; value: %c \n", key, value)
		}
		for {
			time.Sleep(time.Second * 1)
			if i == 8 {
				break
			}
			fmt.Println("i=", i)
			i++
		}

		fmt.Println(strings.Repeat("###", 20))
	}

	{
		var uniques []int
		rand.New(rand.NewSource(time.Now().UnixNano()))
		amount := 5

	loop:
		for len(uniques) < amount {
			randomNum := rand.Intn(amount) + 1
			for _, u := range uniques {
				if u == randomNum {
					fmt.Printf("arr = %v, randomNum = %d is existing, continue loop\n", uniques, randomNum)
					continue loop
				}
			}
			fmt.Printf("arr = %v, randomNum = %d is not existing, appended\n", uniques, randomNum)
			uniques = append(uniques, randomNum)
		}
		sort.Ints(uniques)

		nums := [5]int{5, 4, 3, 2, 1}
		sort.Ints(nums[:])
		fmt.Printf("nums[:] sorted: %v\n", nums)

		fmt.Println(strings.Repeat("###", 20))
	}
}
