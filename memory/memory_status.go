package main

import (
	"fmt"
	"log"
	"runtime"
)

type disappearObj struct {
	a int
	n int
}

func main() {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	fmt.Printf("%d Kb\n", memory.Alloc/1024)

	obj := &disappearObj{1, 2}

	runtime.SetFinalizer(obj, func(newObj *disappearObj) {
		log.Println("test", *newObj)
	})
}
