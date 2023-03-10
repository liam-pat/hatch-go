package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		func() {
			i, j := 10, 20
			defer func() {
				fmt.Printf("i: %d ; j: %d \n", i, j)
			}()
			i, j = i+100, j+100
			fmt.Printf("i: %d ; j: %d \n", i, j)
		}()
		fmt.Println(strings.Repeat("---", 20))
	}
	{
		// defer need to notice case

		// order : result = 0 , defer => result ++ , return
		func() (result int) {
			defer func() {
				result++
			}()
			return
		}()
		// order : result = t , defer => t= t+5 ,return
		func() (result int) {
			t := 5
			defer func() {
				t = t + 5
			}()
			return t
		}()
		// order : r = 0 ,defer => another r = r+5 , return
		func() (r int) {
			defer func(r int) {
				r = r + 5
			}(r)
			return
		}()
	}

	{
		func() {
			i, j := 10, 20
			defer func(i, j int) {
				fmt.Printf("i: %d ; j: %d \n", i, j)
			}(i, j)
			i, j = i+100, j+100
			fmt.Printf("i: %d ; j: %d \n", i, j)
		}()
		fmt.Println(strings.Repeat("---", 20))
	}
	{
		defer func(x int) {
			result := 100 / x
			fmt.Println("result:", result)
		}(0)
		fmt.Println(strings.Repeat("---", 20))
	}

}
