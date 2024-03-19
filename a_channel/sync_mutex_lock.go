package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Container struct {
	mu       sync.Mutex
	userList map[string]int
}

func (c *Container) topUp(name string, amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userList[name] += amount
}

func main() {
	{
		// add the lock to top up
		var wg sync.WaitGroup
		wg.Add(3)

		c := Container{
			userList: map[string]int{"a": 0, "b": 0},
		}

		topUp := func(name string, amount int) {
			c.topUp(name, amount)
			wg.Done()
		}
		go topUp("a", 100)
		go topUp("b", 1000)
		go topUp("a", 10000)

		wg.Wait()
		fmt.Printf("%v\n", c.userList)
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// add lock and no lock
		num1 := 0
		for j := 0; j < 1000; j++ {
			// maybe multiple goroutines get the same value to add , so the result lt 1000
			go func() {
				num1 += 1
			}()
		}
		num2 := 0
		var mutex sync.Mutex
		for i := 0; i < 1000; i++ {
			go func() {
				mutex.Lock()
				defer mutex.Unlock()
				num2 += 1
			}()
		}
		time.Sleep(time.Second * 2)

		fmt.Println("no lock, count = ", num1)
		fmt.Println("add lock, count = ", num2)
	}
}
