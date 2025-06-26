package main

import (
	"fmt"
	"net/http"
	"strings"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	// init get / post params to r.Form
	r.ParseForm()

	fmt.Println("r.Form = ", r.Form)
	fmt.Println("r.URL.Path = ", r.URL.Path)

	for k, v := range r.Form {
		fmt.Printf("%s => %s\n", k, strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello client!")
}

func main() {
	http.HandleFunc("/", indexFunc)

	http.ListenAndServe(":12090", nil)
}
