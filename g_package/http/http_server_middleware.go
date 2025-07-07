package main

import (
	"log"
	"net/http"
)

func access(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("header %s >>>> body %s\n", r.Header, r.Body)
		handle(w, r)
	}
}

func main() {
	http.HandleFunc("/", access(func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("it works")); err != nil {
			log.Println("err: ", err)
			return
		}
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("err: ", err)
		return
	}
}
