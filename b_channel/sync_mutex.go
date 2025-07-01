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

func (c *Container) topup(name string, amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Printf(">> `%s` got the lock to topup %d\n", name, amount)
	c.userList[name] += amount
}

func main() {
	{
		// add the lock to top up
		var wg sync.WaitGroup
		wg.Add(3)

		c := Container{
			userList: map[string]int{"packie": 0, "liam": 0},
		}

		dosth := func(name string, amount int) {
			c.topup(name, amount)
			wg.Done()
		}

		go dosth("packie", 100)
		go dosth("liam", 1000)
		go dosth("packie", 10000)

		wg.Wait()

		fmt.Printf("top up done.... %v \n", c.userList)
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		// add lock and no lock
		num1 := 0
		for j := 0; j < 1000; j++ {
			// maybe multiple goroutines get the same value to add, so the result lt 1000
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

		fmt.Println("no lock to add, count = ", num1)
		fmt.Println("lock to add, count = ", num2)
	}
}
