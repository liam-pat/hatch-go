package main

import "fmt"

func main() {
	noCacheChannel := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			noCacheChannel <- i
		}
		//若 channel 关闭了，就不可以再输入数据进去，但是可以取出数据
		close(noCacheChannel)
	}()

	for true {
		if num, ok := <-noCacheChannel; ok == true {
			fmt.Println("number:", num)
		} else {
			fmt.Println("the channel is close")
			break
		}
	}
}
