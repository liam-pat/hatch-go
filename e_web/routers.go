package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		rootRoute(w, r)
		return
	}
	if len(r.URL.Path) > 6 && r.URL.Path[:6] == "/user/" {
		userRoute(w, r)
		return
	}
	http.NotFound(w, r)
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it works")
}

func userRoute(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Path[6:]
	fmt.Fprintf(w, "your name is %s", userName)
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":12090", mux)
}
