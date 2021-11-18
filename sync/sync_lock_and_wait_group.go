package main

import (
	"fmt"
	"sync"
)

func main() {

	var mt sync.Mutex
	var wg sync.WaitGroup
	var money = 10000

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			mt.Lock()
			fmt.Printf("get %d lock\n", index)
			for j := 0; j < 100; j++ {
				money += 10				//  多个协程对 money产生了竞争
			}
			fmt.Printf("get %d lock\n", index)
			mt.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("money = ", money)		// 应该输出20000才正确
}
