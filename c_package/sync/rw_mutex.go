package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var val = 0

func main() {
	m = new(sync.RWMutex)

	go read()
	go write(2)
	go read()

	time.Sleep(5 * time.Second)
}

func read() {
	m.RLock()
	time.Sleep(1 * time.Second)
	println("val: ", val)
	time.Sleep(1 * time.Second)
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	val = i
	time.Sleep(1 * time.Second)
	m.Unlock()
}
