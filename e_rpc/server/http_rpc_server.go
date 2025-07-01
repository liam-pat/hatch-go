package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserService struct{}

func (u *UserService) Desc(Args *User, reply *string) error {
	*reply = fmt.Sprintf("from server: name: %s, age: %d", Args.Name, Args.Age)
	return nil
}

func main() {
	rpc.Register(new(UserService))
	rpc.HandleHTTP()
	_ = http.ListenAndServe(":1234", nil)
}
