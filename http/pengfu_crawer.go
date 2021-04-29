package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func spiderListContent(page int) (content string, err error) {
	url := fmt.Sprintf("http://www.pengfu.com/xiaohua_%s.html", strconv.Itoa(page))
	fmt.Printf("Crawling page %d , url : %s \n", page, url)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println("http get err", err)
		return
	}
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

	return
}

func spiderDetail(url string) (title, articleContent string, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err2 := client.Get(url)
	if err2 != nil {
		err = err2
		return
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			err = err1
			return
		}

	}(res.Body)

	buf := make([]byte, 4*1024)
	var content string
	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[:n])
	}

	// get title
	re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re == nil {
		err = fmt.Errorf("regexp,MustCompile err")
		return
	}
	tmpTitles := re.FindAllStringSubmatch(content, 1)
	for _, data := range tmpTitles {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	// get article
	re1 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re1 == nil {
		err = fmt.Errorf("regexp,MustCompile err")
		return
	}
	tmpArticles := re1.FindAllStringSubmatch(content, -1)
	for _, data := range tmpArticles {
		articleContent = data[1]
		articleContent = strings.Replace(articleContent, "\t", "", -1)
		articleContent = strings.Replace(articleContent, "\n", "", -1)
		articleContent = strings.Replace(articleContent, "\r", "", -1)

		break
	}

	return
}

func storeJoyToFile(i int, fileTitle, fileContent []string) {
	f, err := os.Create("./http/down_file/joy_" + strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("open file got error", err)
		return
	}
	defer f.Close()

	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")
		f.WriteString(fileContent[i] + "\n")
		f.WriteString("\n-----------------------")
	}
}

func doWork(start, end int) {
	fmt.Printf("Start to spider page %d to page %d \n", start, end)

	finishPage := make(chan int)

	for i := start; i <= end; i++ {
		go func(i int, finishPage chan int) {
			content, _ := spiderListContent(i)
			re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
			if re == nil {
				fmt.Println("regexp.MustCompile err")
				return
			}
			fileTitle := make([]string, 0)
			fileArticle := make([]string, 0)
			matchUrls := re.FindAllStringSubmatch(content, -1)
			for _, data := range matchUrls {
				title, article, err := spiderDetail(data[1])
				if err != nil {
					fmt.Println("spider detail err", err)
					continue
				}
				fileTitle = append(fileTitle, title)
				fileArticle = append(fileArticle, article)
			}

			fmt.Println("################################")
			fmt.Printf("page : %d content :", i)
			fmt.Println(fileTitle, fileArticle)
			fmt.Println("################################")
			storeJoyToFile(i, fileTitle, fileArticle)

			finishPage <- i
		}(i, finishPage)
	}

	for j := start; j <= end; j++ {
		fmt.Printf("-------------------------------------page %d already finish\n", <-finishPage)
	}
}

func main() {
	var start, end int

	fmt.Printf("input start page :")
	fmt.Scan(&start)
	fmt.Printf("input end page :")
	fmt.Scan(&end)

	doWork(start, end)

}
