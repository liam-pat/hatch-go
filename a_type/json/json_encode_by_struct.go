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
	{
		str := data{"pccw", "packie", []string{"boy", "26"}, 11111}

		buf, err := json.Marshal(str)

		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("format json :", string(buf))
	}
	{
		jsonStr := `{"company":"pccw","name":"packie","detail":["boy","26"],"price":11111}`
		var tmp data
		err := json.Unmarshal([]byte(jsonStr), &tmp)

		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("tmp = %+v \n", tmp)
	}

}
