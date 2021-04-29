package main

import (
	"fmt"
	"net/http"
)

func HandConn(res http.ResponseWriter, req *http.Request) {
	fmt.Println("r.method", req.Method)

	_, err := res.Write([]byte("hello world"))
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/", HandConn)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
