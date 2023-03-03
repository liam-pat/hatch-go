package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("http://127.0.0.1:8000")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	fmt.Printf("status = %v \n header = %v \n statusCode = %v \n", res.Status, res.Header, res.StatusCode)

	buf := make([]byte, 1024)
	var tmp string

	for {
		n, _ := res.Body.Read(buf)
		if n == 0 {
			break
		}
		tmp += string(buf[:n])
	}

	fmt.Println("body = ", tmp)
}
