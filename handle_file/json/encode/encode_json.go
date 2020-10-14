package main

import (
	"encoding/json"
	"fmt"
)

//notice the param should upCase
type car struct {
	Name   string   `json:"car_name"`
	Colors []string `json:"car_colors"`
	Size   int      `json:"-"`
}

func main() {
	newCar := car{"tokyo", []string{"red", "blue", "yellow"}, 12}
	//byteArr, err := json.Marshal(newCar)
	byteArr, err := json.MarshalIndent(newCar, "", " ")

	if err != nil {
		fmt.Println("have something error")
		return
	} else {
		fmt.Println(string(byteArr))
	}

	//cover map to json
	mapArr := make(map[string]interface{}, 4)
	mapArr["name"] = "bil"
	mapArr["age"] = 12
	mapArr["sex"] = "boy"
	mapArr["job"] = []string{"police", "coder", "teacher", "student"}

	byteMaps, err := json.Marshal(mapArr)

	if err != nil {
		println(err)
		return
	} else {
		println(string(byteMaps))
	}
}
