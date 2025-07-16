package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	{
		var waitGroup sync.WaitGroup
		var urls = []string{"https://www.baidu.com", "https://www.163.com", "https://www.weibo.com"}

		for _, url := range urls {
			waitGroup.Add(1)
			go func(url string) {
				defer waitGroup.Done()
				_, err := http.Get(url)
				fmt.Println("request: ", url, err)
			}(url)
		}

		waitGroup.Wait()
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		var mutex sync.Mutex
		var waitGroup sync.WaitGroup
		var money = 10000

		for i := 0; i < 5; i++ {
			waitGroup.Add(1)
			go func(index int) {
				mutex.Lock()
				fmt.Printf("%d got lock\n", index)
				for j := 0; j < 100; j++ {
					money += 10 //  race condition
				}
				mutex.Unlock()
				waitGroup.Done()
			}(i)
		}
		waitGroup.Wait()

		fmt.Println("money = ", money)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		var once sync.Once
		var waitGroup sync.WaitGroup

		p := 0

		for i := 0; i < 5; i++ {
			waitGroup.Add(1)
			go func() {
				once.Do(func() { p += 1 })
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()

		fmt.Println("age ：", p)
		fmt.Println(strings.Repeat("##", 20))
	}
}
