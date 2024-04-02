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
		str := data{
			Company:   "IBM",
			Name:      "liam",
			Detail:    []string{"boy", "26"},
			Price:     11111,
			OtherName: "",
			Amount:    100,
		}

		if buf, err := json.Marshal(str); err != nil {
			fmt.Println("err = ", err)
			return
		} else {
			fmt.Println("format json :", string(buf))
		}

		fmt.Println(strings.Repeat("##", 20))
	}
	{
		jsonStr := `{"company":"ibm","name":"liam","detail":["boy","26"],"price":11111}`
		var tmp data
		if err := json.Unmarshal([]byte(jsonStr), &tmp); err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("tmp = %+v \n", tmp)

		fmt.Println(strings.Repeat("##", 20))

	}
}
