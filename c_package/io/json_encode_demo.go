package main

import (
	"encoding/json"
	"fmt"
)

type car struct {
	Name   string   `json:"car_name"`
	Colors []string `json:"car_colors"`
	Size   int      `json:"-"`
}

func main() {
	newCar := car{
		"tokyo",
		[]string{"red", "blue", "yellow"},
		12,
	}
	byteArr, _ := json.MarshalIndent(newCar, "", " ")
	fmt.Println(string(byteArr))

	mapArr := make(map[string]interface{}, 4)
	mapArr["name"] = "bil"
	mapArr["age"] = 12
	mapArr["sex"] = "boy"
	mapArr["job"] = []string{"police", "coder", "teacher", "student"}

	byteMaps, _ := json.Marshal(mapArr)
	fmt.Println(string(byteMaps))
}
