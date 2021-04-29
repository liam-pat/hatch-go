package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func spiderPage(page int, doChan chan<- int) {
	initUrl := "https://tieba.baidu.com/f/good?kw=%E5%87%A1%E4%BA%BA%E4%BF%AE%E4%BB%99%E4%BC%A0&ie=utf-8&cid=0"
	url := fmt.Sprintf("%s&pn=%d", initUrl, (page-1)*50)
	fmt.Printf("start to crawl page %d , url = %s\n", page, url)

	pageContent, _ := getPageContent(url)

	dir := "./http/down_file/"
	_ = os.Mkdir(dir, 0775)
	fileName := dir + strconv.Itoa(page) + ".html"
	fileSign, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("create file error", err1)
		return
	}

	_, err := fileSign.Write([]byte(pageContent))
	if err != nil {
		fmt.Println("write to file error", err)
		return
	}

	doChan <- page
}

func getPageContent(url string) (content string, err error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get page content error", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("close http body error", err)
		}
	}(res.Body)

	buffer := make([]byte, 1024*4)
	for {
		n, err1 := res.Body.Read(buffer)
		// maybe EOF is coming
		if n == 0 {
			err = err1
			break
		}
		content += string(buffer[:n])
	}

	return
}

func do(start, end int) {
	fmt.Printf("start to crawl page %d to page %d\n", start, end)
	doChan := make(chan int)
	for page := start; page <= end; page++ {
		go spiderPage(page, doChan)
	}

	for {
		select {
		case page := <-doChan:
			fmt.Printf("page %d already finish\n", page)
		case <-time.After(10 * time.Second):
			fmt.Println("10s go away, stop to run")
			return
		}
	}
}

func main() {
	var start, end int
	fmt.Printf("input the start page : ")
	_, err := fmt.Scan(&start)
	if err != nil {
		return
	}

	fmt.Printf("input the end page : ")
	_, err1 := fmt.Scan(&end)
	if err1 != nil {
		return
	}

	do(start, end)
}
