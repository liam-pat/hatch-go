package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	if client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234"); err != nil {
		log.Fatal("dialing:", err)
	} else {
		var reply string
		user := User{Name: "liam", Age: 40}

		err = client.Call("UserService.Desc", &user, &reply)
		if err != nil {
			log.Fatal("err", err)
		}
		fmt.Printf("%s\n", reply)
	}
}
