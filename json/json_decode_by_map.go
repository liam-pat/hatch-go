package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonStr := ` {"company":"pccw","detail":["boy","26"],"name":"packie","price":"111111"}`

	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonStr), &m)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Printf("m = %+v\n", m)

}
