package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"https://www.baidu.com/",
		"https://www.163.com/",
		"https://www.weibo.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}

	wg.Wait()
	fmt.Println("over")
}
