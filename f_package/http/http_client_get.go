package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.baidu.com")

	defer func(Body io.ReadCloser) { Body.Close() }(res.Body)

	fmt.Printf("status = %v \nheader = %v \nstatusCode = %v \n", res.Status, res.Header, res.StatusCode)

	buf := make([]byte, 2048)
	var contentStr string
	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		contentStr += string(buf[:n])
	}
	fmt.Println("body = ", contentStr)
}
