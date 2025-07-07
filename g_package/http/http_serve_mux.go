package main

import (
	"log"
	"net/http"
)

func main() {
	userMux := http.NewServeMux()
	adminMux := http.NewServeMux()

	userMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("request /user ")); err != nil {
			log.Println("err: ", err)
			return
		}
	})
	adminMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("request /admin")); err != nil {
			log.Println("err: ", err)
			return
		}
	})
	userServer := http.Server{Addr: ":8080", Handler: userMux}
	adminServer := http.Server{Addr: ":8081", Handler: adminMux}
	if err := userServer.ListenAndServe(); err != nil {
		return
	}
	if err := adminServer.ListenAndServe(); err != nil {
		return
	}
}
