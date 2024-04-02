package main

import (
	"fmt"
	"runtime"
)

func main() {

	// run multiple cores by default after go1.5
	cores := runtime.NumCPU()            // number of the cpu core
	threads := runtime.GOMAXPROCS(cores) // number of the thread

	fmt.Println("cpu num: ", cores)
	fmt.Println("thread num: ", threads)
}
