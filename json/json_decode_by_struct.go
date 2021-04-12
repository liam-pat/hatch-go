package main

import (
	"encoding/json"
	"fmt"
)

type data01 struct {
	Company string   `json:"company"`
	Name    string   `json:"name"`
	Detail  []string `json:"detail"`
	Price   int      `json:"-"`
}

func main() {
	jsonStr := `{"company":"pccw","name":"packie","detail":["boy","26"],"price":11111}`

	var tmp data01
	err := json.Unmarshal([]byte(jsonStr), &tmp)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp = %+v \n", tmp)
}
