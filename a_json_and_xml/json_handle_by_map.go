package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	{
		data := make(map[string]interface{}, 4)
		data["company"] = "ibm"
		data["name"] = "liam"
		data["price"] = "111111"
		data["detail"] = []string{"boy", "26"}

		if buf, err := json.Marshal(data); err != nil {
			fmt.Println("err = ", err)
			return
		} else {
			fmt.Println("format json :", string(buf))
		}
	}

	{
		jsonStr := ` {"company":"ibm","detail":["boy","26"],"name":"liam","price":"111111"}`
		m := make(map[string]interface{}, 4)
		if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
			fmt.Println("err", err)
			return
		}

		fmt.Printf("m = %+v\n", m)
	}
}
