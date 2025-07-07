package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func spiderPage(page int, doChan chan<- int) {

	url := "https://tieba.baidu.com/f/good?kw=%E5%87%A1%E4%BA%BA%E4%BF%AE%E4%BB%99%E4%BC%A0&ie=utf-8&cid=0"
	url = fmt.Sprintf("%s&pn=%d", url, (page-1)*50)

	content, _ := getContent(url)
	fmt.Println("content: ", content)
	doChan <- page
}

func getContent(url string) (content string, err error) {
	res, _ := http.Get(url)
	defer func(Body io.ReadCloser) { _ = Body.Close() }(res.Body)

	buffer := make([]byte, 1024*4)
	for {
		n, err := res.Body.Read(buffer)
		if n == 0 || err == io.EOF {
			break
		}
		content += string(buffer[:n])
	}
	return
}

func do(start, end int) {
	doChan := make(chan int)
	for page := start; page <= end; page++ {
		go spiderPage(page, doChan)
	}
	for {
		select {
		case page := <-doChan:
			fmt.Printf("page %d finished \n", page)
		case <-time.After(5 * time.Second):
			return
		}
	}
}

func main() {
	var start, end int

	fmt.Printf("input the start page : ")
	_, _ = fmt.Scan(&start)
	fmt.Printf("input the end page : ")
	_, _ = fmt.Scan(&end)

	do(start, end)
}
