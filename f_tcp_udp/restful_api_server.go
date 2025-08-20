package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprintln(res, "it works!")
}

func Get(res http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(res, "%s welcome\n", params.ByName("name"))
}

func Rename(res http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(res, "rename to %s\n", params.ByName("name"))
}

func Delete(res http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(res, "%s deleted\n", params.ByName("name"))
}

func Add(res http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(res, "%s added! \n", params.ByName("name"))
}

func main() {
	/**
	eg.
	curl http://127.0.0.1:8999/user/dog
	curl -X POST http://127.0.0.1:8999/user/dog
	*/
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/user/:name", Get)

	router.POST("/user/:name", Add)
	router.DELETE("/user/:name", Delete)
	router.PUT("/user/:name", Rename)

	log.Fatal(http.ListenAndServe(":8999", router))
}
