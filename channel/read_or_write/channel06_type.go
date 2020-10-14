package main

import "fmt"

//only write ,channel 传参是引用传递
func producer(canIn chan<- int) {
	for i := 0; i < 4; i++ {
		canIn <- i * i
	}
	close(canIn)
}

//only read
func consumer(canOut <-chan int) {
	for num := range canOut {
		fmt.Println("num = ", num)
	}
}

func main() {
	//双向的 channel
	channelExample := make(chan int)

	//双向 channel 隐式转换成单向 channel ,
	//but 单向的 channel 无法转换成双向的

	//var writeChannel chan<- int = channelExample // only write
	//var readChannel <-chan int = channelExample  // only read

	go producer(channelExample)
	consumer(channelExample)
}
