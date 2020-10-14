package main

import (
	"fmt"
	"time"
)

//协程 coroutine


func main() {
	//其他语言的一般进程开到100个基本上是资源上限了但是 golang 引用了协程的概念，将协程开到1000也没有很大的问题
	for i := 0; i < 100; i++ {
		go func(i int) {
			for true {
				fmt.Printf("Hello form goroutine %d \n", i)
			}
		}(i)
	}
	//若没有主进程的等待，协程还没有创建成功就结束了
	time.Sleep(time.Millisecond)
}
