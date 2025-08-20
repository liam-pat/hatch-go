package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func spiderContent(url string, page int) (content string, err error) {
	if page > 1 {
		url = fmt.Sprintf("%s&page=%d&o=20", url, page)
	}

	fmt.Printf("Start get page %d  ----->  %s \n", page, url)

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	res, _ := client.Do(req)
	defer func(Body io.ReadCloser) { _ = Body.Close() }(res.Body)

	buf := make([]byte, 10*1024)
	for {
		n, err := res.Body.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		content += string(buf[:n])
	}
	fmt.Println("content: ", string(buf))
	return
}

func doWork(url string, start, end int) {
	finish := make(chan int)
	for page := start; page <= end; page++ {
		go func(i int, finishPage chan int) {
			spiderContent(url, page)
			finishPage <- page
		}(page, finish)
	}

	for {
		select {
		case page := <-finish:
			fmt.Printf("page %d finished \n", page)
		case <-time.After(5 * time.Second):
			return
		}
	}
}

func main() {
	var start, end int
	url := "https://search.bilibili.com/all?keyword=goland&search_source=1"

	fmt.Printf("start page :")
	fmt.Scan(&start)
	fmt.Printf("end page :")
	fmt.Scan(&end)

	doWork(url, start, end)
}
