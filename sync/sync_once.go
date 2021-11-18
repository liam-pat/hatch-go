package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grown() {
	p.Age += 1
}

func main() {

	var once sync.Once
	var wg sync.WaitGroup

	p := &Person{
		"bone",
		0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() { p.Grown() })
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("age ：", p.Age)
}
