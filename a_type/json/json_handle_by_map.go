package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	{
		data := make(map[string]interface{}, 4)

		data["company"] = "pccw"
		data["name"] = "packie"
		data["price"] = "111111"
		data["detail"] = []string{"boy", "26"}

		buf, err := json.Marshal(data)

		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("format json :", string(buf))
	}
	{
		jsonStr := ` {"company":"pccw","detail":["boy","26"],"name":"packie","price":"111111"}`
		m := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(jsonStr), &m)

		if err != nil {
			fmt.Println("err", err)
		}

		fmt.Printf("m = %+v\n", m)
	}
}
