package main

import "fmt"

func main() {

	{
		var num int
		fmt.Printf("first area, input your number: ")
		fmt.Scan(&num)

		// `break` auto append ,
		// if has `fallthrough` the code will go to next case

		switch num {
		case 1:
			fmt.Printf("num: %d \n", num)
		case 2:
			fmt.Printf("num: %d \n", num)
		case 4:
			fmt.Printf("num: %d \n", num)
		case 5:
			fmt.Printf("num: %d \n", num)

		default:
			fmt.Println("do not have something to match")
		}
	}
	{
		switch sample := 1; sample {
		case 1, 2, 3:
			fmt.Println("match")
		default:
			fmt.Println("do not match")
		}
	}
	{
		var score int
		fmt.Printf("second area, input your score: ")
		fmt.Scan(&score)

		switch true {
		case score > 90:
			fmt.Println("prefect")
		case score > 80:
			fmt.Println("nice")
		case score < 60:
			fmt.Println("continue")
		default:
			fmt.Println("keep it up!")
		}
	}
}
