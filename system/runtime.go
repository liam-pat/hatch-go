package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	fmt.Printf("PC has %d core",runtime.NumCPU())

	// using more core
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 1; i <= 10; i++ {
		go func(i int){
			if i == 5 {
				runtime.Gosched()	// never output 5 firstly
				//runtime.Goexit()  // exit this goroutine
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
