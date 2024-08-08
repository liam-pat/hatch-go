package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(res, "Welcome!\n")
}

func GetUser(res http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	_, _ = fmt.Fprintf(res, "you are geting user %s\n", name)
}

func ModifyUser(res http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	_, _ = fmt.Fprintf(res, "you are modifing user %s\n", name)
}

func DeleteUser(res http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	_, _ = fmt.Fprintf(res, "you are deleting user %s\n", name)
}

func AddUser(res http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	_, _ = fmt.Fprintf(res, "you are adding user %s\n", name)
}

func main() {
	/**
	eg.
	curl http://127.0.0.1:8999/user/dog
	curl -X POST http://127.0.0.1:8999/user/dog
	*/
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/user/:name", GetUser)

	router.POST("/user/:name", AddUser)
	router.DELETE("/user/:name", DeleteUser)
	router.PUT("/user/:name", ModifyUser)

	log.Fatal(http.ListenAndServe(":8999", router))
}
