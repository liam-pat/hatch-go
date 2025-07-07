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
	// if the second parameter is nil, the default http.DefaultServeMux will be used, http.DefaultServeMux.HandleFunc("/", myHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("err: ", err)
		return
	}
}
