package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOkay   bool     `json:"is_okay"`
	Price    float64  `json:"price"`
}

func main() {
	jsonString := `
{
	"company":"lv",
	"subjects":[
		"Go",
		"C++",
		"Python",
		"Test"
	],
	"is_okay": true,
	"price": 666.666
}`
	var structIT IT
	json.Unmarshal([]byte(jsonString), &structIT)
	fmt.Printf("structIT:%+v\n", structIT)

	mapSample := make(map[string]interface{}, 4)
	json.Unmarshal([]byte(jsonString), &mapSample)

	//转换成 map 类型的时候，因为 value type =  interface ，所有要类型断言
	for key, value := range mapSample {
		switch value.(type) {

		case string:
			fmt.Printf("map[%s] type is string , value = %s \n", key, value)
		case bool:
			fmt.Printf("map[%s] type is bool , value = %v \n", key, value)
		case float64:
			fmt.Printf("map[%s] type is float64 , value = %f \n", key, value)
		case []interface{}:
			fmt.Printf("map[%s] type is []interface{} , value = %v \n", key, value)
		}
	}
}
