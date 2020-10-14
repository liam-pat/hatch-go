package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

/**
首先是先输出 Hello, 当第二个协程从 channel 取数据的时候，因为person1 还没有输出完 hello，所以会进入阻塞
当输出完 Hello,往 channel 里面输入 1 ，那么第二个进程检查到 channel 里面有数据，就会解除阻塞,输出 word
 */
func person1() {
	printer("hello")
	channel <- 666
}
func person2() {
	<-channel
	printer("world")
}

func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func main() {
	go person1()
	go person2()
	for ; ;  {
		
	}
}
