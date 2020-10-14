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
	err := json.Unmarshal([]byte(jsonString), &structIT)

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("structIT:%+v\n", structIT)




	mapSample := make(map[string]interface{}, 4)
	err = json.Unmarshal([]byte(jsonString), &mapSample)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//转换成 map 类型的时候，因为 key 是万能类型 interface ，所有要类型断言
	for key, value := range mapSample {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s] type is string , value = %s \n", key, data)
		case bool:
			fmt.Printf("map[%s] type is bool , value = %v \n", key, data)
		case float64:
			fmt.Printf("map[%s] type is float64 , value = %f \n", key, data)
		case []interface{}:
			fmt.Printf("map[%s] type is []interface{} , value = %v \n", key, data)
		}
	}
}
