package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("r.Form = ", r.Form)
	fmt.Println("r.URL.Path = ", r.URL.Path)
	fmt.Println("r.URL.Schem = ", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Printf("key: %s -- value %s\n", k, strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello client!")
}

func main() {
	fmt.Println("testing")

	http.HandleFunc("/", indexFunc)
	err := http.ListenAndServe(":12090", nil)

	if err != nil {
		log.Fatal("Listen And Serve: ", err)
	}
}
