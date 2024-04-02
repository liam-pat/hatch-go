package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("request /")); err != nil {
			log.Println("err: ", err)
			return
		}
	})
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("r.method", request.Method)
		if _, err := writer.Write([]byte("request /hello")); err != nil {
			log.Println("err: ", err)
			return
		}
	})
	mux.HandleFunc("/world", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("r.method", request.Method)
		if _, err := writer.Write([]byte("request /world")); err != nil {
			log.Println("err: ", err)
			return
		}
	})

	server := http.Server{Addr: ":8080", Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		return
	}
}
