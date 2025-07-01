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
				fmt.Printf("%-5s >> i: %d ; j: %d \n", "defer", i, j)
			}()
			i, j = i+100, j+100
			fmt.Printf("%-5s >> i: %d ; j: %d \n", "main", i, j)
		}()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		// order : result = 0 , defer => result ++ , return
		func() (result int) {
			defer func() { result++ }()
			return
		}()
		// order : result = t , defer => t= t+5 ,return
		func() (result int) {
			t := 5
			defer func() { t = t + 5 }()
			return t
		}()
		// order : r = 0 ,defer => another r = r+5 , return
		func() (result int) {
			defer func(result int) { result = result + 5 }(result)
			return
		}()
		fmt.Println(strings.Repeat("##", 20))
	}
}
