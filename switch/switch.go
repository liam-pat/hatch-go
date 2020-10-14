package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Please input your floor : ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Printf("num: %d \n", num)
		//default have break
	case 2:
		fmt.Printf("num: %d \n", num)
	case 4:
		fmt.Printf("num: %d \n", num)
		//you can use fallthrough to go on
	case 5:
		fmt.Printf("num: %d \n", num)

	default:
		fmt.Println("do not have something to match");
	}

	//you also can use that
	switch sample := 1; sample {
	case 1, 2, 3:
		fmt.Println("match")
	default:
		fmt.Println("do not match")
	}

	var score int
	fmt.Printf("Please input your point : ")
	fmt.Scan(&score)

	switch true {
	case score > 90:
		fmt.Println("prefect point")
	case score > 80:
		fmt.Println("nice point")
	case score < 60:
		fmt.Println("continue")
	default:
		fmt.Println("not good")
	}
}
