package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type data struct {
	Company   string   `json:"company"`
	Name      string   `json:"name"`
	Detail    []string `json:"detail"`
	Price     int      `json:"-"`                    // does not display
	OtherName string   `json:"other_name,omitempty"` // keyword `omitempty` means does not echo when value = ""
	Amount    int      `json:"amount,string"`        // covert `bool, string, int, int64` to string
}

func main() {
	{
		data := make(map[string]interface{}, 4)

		data["company"] = "ibm"
		data["name"] = "liam"
		data["price"] = "111111"
		data["detail"] = []byte("If I were a boy~")

		if buf, err := json.Marshal(data); err != nil {
			fmt.Println("err = ", err)
			return
		} else {
			fmt.Println("json :", string(buf))
		}
	}
	{
		jsonstr := `{"company":"ibm","detail":["boy","26"],"name":"liam","price":"111111"}`

		m := make(map[string]interface{}, 4)

		if err := json.Unmarshal([]byte(jsonstr), &m); err != nil {
			fmt.Println("err", err)
			return
		}
		fmt.Printf("m: %+v\n", m)
	}
	{
		info := data{Company: "IBM", Name: "liam", Detail: []string{"boy", "26"}, Price: 11111, OtherName: "", Amount: 100}

		if buf, err := json.Marshal(info); err != nil {
			fmt.Println("err: ", err)
			return
		} else {
			fmt.Println("json: ", string(buf))
		}

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		jsonStr := `{"company":"ibm","name":"liam","detail":["boy","26"],"price":11111}`
		var data data

		if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("tmp:%+v \n", data)
	}
}
