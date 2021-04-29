package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http get err =", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	fmt.Println("status = ", res.Status)
	fmt.Println("header = ", res.Header)
	fmt.Println("statusCode = ", res.StatusCode)

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
