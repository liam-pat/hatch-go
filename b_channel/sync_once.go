package main

import (
	"fmt"
	"sync"
)

type person struct {
	Name string
	Age  int
}

func (p *person) grown() {
	p.Age += 1
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	p := &person{
		"liam",
		0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() { p.grown() }) // only exec one time
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("age is：", p.Age)
}
