package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	Company string   `json:"company"`
	Name    string   `json:"name"`
	Detail  []string `json:"detail"`
	Price   int      `json:"-"`
}

func main() {
	str := data{"pccw", "packie", []string{"boy", "26"}, 11111}

	buf, err := json.Marshal(str)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("format json :", string(buf))
}
