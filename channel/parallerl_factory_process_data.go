package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int, 100)

	for i := 0; i < 5; i++ {
		in <- i
		go parallelProcessData(in)
	}

	time.Sleep(10 * time.Second)
}

func parallelProcessData(in <-chan int) {
	// make channels:
	preOut := make(chan int, 100)
	stepAOut := make(chan int, 100)
	stepBOut := make(chan int, 100)
	stepCOut := make(chan int, 100)
	out := make(chan int, 100)
	// start parallel computations:
	go preprocessData(in, preOut)
	go processStepA(preOut, stepAOut)
	go processStepB(stepAOut, stepBOut)
	go processStepC(stepBOut, stepCOut)
	go postProcessData(stepCOut, out)

	if <-out == 1 {
		fmt.Println("####one item is finish####")
	}
}

func preprocessData(in <-chan int, out chan<- int) {
	getOne := <-in
	fmt.Println("1. Process Data, channel data:", getOne)
	out <- 1
}

func processStepA(in <-chan int, out chan<- int) {
	getOne := <-in
	fmt.Println("2. Process Data Step A, channel data:", getOne)
	out <- 1
}

func processStepB(in <-chan int, out chan<- int) {
	getOne := <-in
	fmt.Println("3. Process Data Step B, channel data:", getOne)
	out <- 1
}

func processStepC(in <-chan int, out chan<- int) {
	getOne := <-in
	fmt.Println("4. Process Data Step C, channel data:", getOne)
	out <- 1
}

func postProcessData(in <-chan int, out chan<- int) {
	getOne := <-in
	fmt.Println("5. Post Data Step C,channel data:", getOne)
	out <- 1
}
