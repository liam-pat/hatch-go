package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		root(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "okay, you enterred")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":12090", mux)
}
