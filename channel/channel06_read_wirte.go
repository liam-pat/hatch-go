package main

import "fmt"

func product(writer chan<- int) {
	for i := 0; i < 10; i++ {
		writer <- i * i
	}
	close(writer)
}

func costumer(reader <-chan int) {
	for num := range reader {
		fmt.Println("reader the number", num)
	}
}

func main() {
	// create the no cache channel also can write and read
	ch := make(chan int)

	// start to product data
	go product(ch)

	// start to get the data
	costumer(ch)
}
