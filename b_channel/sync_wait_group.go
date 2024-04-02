package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	{
		// request multiple urls
		var wg sync.WaitGroup

		var urls = []string{"https://www.baidu.com", "https://www.163.com", "https://www.weibo.com"}
		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				res, err := http.Get(url)
				fmt.Printf("url %s ,%v , has err %v\n", url, res.Header, err)
			}(url)
		}
		wg.Wait()
		fmt.Println(strings.Repeat("##", 30))
	}
	{
		var mt sync.Mutex
		var wg sync.WaitGroup
		var money = 10000

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(index int) {
				mt.Lock()
				defer mt.Unlock()
				fmt.Printf("goroutine %d get lock\n", index)
				for j := 0; j < 100; j++ {
					money += 10
				}
				fmt.Printf("goroutine %d unlock\n", index)
				wg.Done()
			}(i)
		}
		wg.Wait()
		fmt.Println("amount :  ", money)
	}
}
