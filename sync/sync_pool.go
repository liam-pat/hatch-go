package main

import (
	"fmt"
	"sync"
	"time"
)

var BytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {

	time1 := time.Now().Unix()
	obj1 := make([]byte, 1024)
	for i := 0; i < 100000000; i++ {
		obj1 = make([]byte, 1024)
		_ = obj1
	}

	time2 := time.Now().Unix()
	obj2 := BytePool.Get().(*[]byte)
	for i := 0; i < 100000000; i++ {
		obj2 = BytePool.Get().(*[]byte)
		BytePool.Put(obj2)
		_ = obj2
	}

	time3 := time.Now().Unix()

	fmt.Println("without pool: ", time2-time1, "s") // 16s
	fmt.Println("with    pool: ", time3-time2, "s") // 1s
}
