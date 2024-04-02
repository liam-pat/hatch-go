package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("r.method", request.Method)
		if _, err := writer.Write([]byte("hello world")); err != nil {
			log.Println("err: ", err)
			return
		}
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("err: ", err)
		return
	}
}
