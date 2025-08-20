package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type disappearObj struct {
	a int
	n int
}

func main() {
	{
		var memory runtime.MemStats

		runtime.ReadMemStats(&memory)
		fmt.Printf("used %d Kb\n", memory.Alloc>>10)

		obj := &disappearObj{1, 2}

		runtime.SetFinalizer(obj, func(newObj *disappearObj) {
			log.Println("the object is disappear", *newObj)
		})

		obj = nil
		runtime.GC()

		time.Sleep(time.Second)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		fmt.Printf("PC has %d core \n", runtime.NumCPU())
		runtime.GOMAXPROCS(runtime.NumCPU())

		for i := 1; i <= 10; i++ {
			go func(i int) {
				if i == 5 {
					runtime.Gosched() // never output 5 firstly
					//runtime.Goexit()  // exit this goroutine
				}
				fmt.Println("goroutine ", i)
			}(i)
		}
		time.Sleep(time.Second)

		fmt.Println(strings.Repeat("##", 20))
	}
}
