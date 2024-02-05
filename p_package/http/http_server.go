package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("r.method", request.Method)

		writer.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8000", nil)
}
