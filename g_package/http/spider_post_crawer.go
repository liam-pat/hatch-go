package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func spiderPage(page int, doChan chan<- int) {

	dir := "./files/"
	_ = os.Mkdir(dir, 0775)

	initUrl := "https://tieba.baidu.com/f/good?kw=%E5%87%A1%E4%BA%BA%E4%BF%AE%E4%BB%99%E4%BC%A0&ie=utf-8&cid=0"

	url := fmt.Sprintf("%s&pn=%d", initUrl, (page-1)*50)
	fmt.Printf("---start goruntine to crawl page %d ----> %s\n", page, url)

	pageContent, _ := getPageContent(url)

	fileName := dir + strconv.Itoa(page) + ".html"
	file, _ := os.Create(fileName)

	file.Write([]byte(pageContent))

	doChan <- page
}

func getPageContent(url string) (content string, err error) {
	res, _ := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("close http body error", err)
		}
	}(res.Body)

	buffer := make([]byte, 1024*4)
	for {
		n, _ := res.Body.Read(buffer)
		if n == 0 {
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
			fmt.Printf("page %d has already finished \n", page)
		case <-time.After(10 * time.Second):
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

	fmt.Println(strings.Repeat("#*", 20))

	do(start, end)
}
