package main

import (
	"log"
	"net/http"
)

func accessLog(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("header %s >>>> body %s\n", r.Header, r.Body)
		handle(w, r)
	}
}

func main() {

	http.HandleFunc("/", accessLog(func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("success")); err != nil {
			log.Println("err: ", err)
			return
		}
	}))

	server := http.Server{Addr: "127.0.0.1:8080"}

	if err := server.ListenAndServe(); err != nil {
		log.Println("err: ", err)
		return
	}
}
