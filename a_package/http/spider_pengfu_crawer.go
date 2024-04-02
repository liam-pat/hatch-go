package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func spiderListContent(page int) (content string, err error) {
	pageUrl := "https://search.bilibili.com/all?keyword=test&search_source=1"
	if page > 1 {
		pageUrl = fmt.Sprintf("https://search.bilibili.com/all?keyword=test&search_source=1&page=%d&o=20", page)
	}

	fmt.Printf("Start get page %d  ----->  %s \n", page, pageUrl)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, _ := client.Get(pageUrl)
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("http get err", err1)
		}
	}(res.Body)

	buf := make([]byte, 4*1024)

	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[:n])
	}
	file, _ := os.Create(fmt.Sprintf("./files/bilibil_%d.html", page))

	file.Write([]byte(content))

	return
}

func doWork(start, end int) {
	finishPage := make(chan int)

	for i := start; i <= end; i++ {
		go func(i int, finishPage chan int) {
			spiderListContent(i)
			finishPage <- i
		}(i, finishPage)
	}

	for {
		select {
		case page := <-finishPage:
			fmt.Printf("page %d has already finished \n", page)
		case <-time.After(10 * time.Second):
			return
		}
	}
}

func main() {
	var start, end int

	fmt.Printf("start page :")
	fmt.Scan(&start)
	fmt.Printf("end page :")
	fmt.Scan(&end)

	doWork(start, end)

}
