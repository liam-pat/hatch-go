package main

import "fmt"

// init the map type
// map[keyType]valueType
// map 打印出来可能是乱序的 ,map 只有 len 没有 cap，

func main() {
	var mapSample1 map[string]string
	fmt.Println("init the empty mapSample1: ", mapSample1)
	fmt.Printf("mapSample2 Length: %d \n", len(mapSample1))
	fmt.Println("--------------------------------------------")

	fmt.Println("it also can init by make")
	//you also can init the length by make
	mapSample2 := make(map[string]string, 2)
	mapSample2["firstLesson"] = "Java"
	mapSample2["secondLesson"] = "C Plus Plus"
	mapSample2["thirdLesson"] = "PHP"
	fmt.Println("init the  mapSample2 by make: ", mapSample2)
	fmt.Printf("mapSample2 length: %d \n", len(mapSample2))
	fmt.Println("--------------------------------------------")

	fmt.Println("print the map by for")
	for key, value := range mapSample2 {
		fmt.Printf("key: %s \t value: %s \n", key, value)
	}
	fmt.Println("--------------------------------------------")

	// mapSample2 mean
	checkKey := "firstLesson"
	value, isExist := mapSample2[checkKey]

	if isExist {
		fmt.Printf("key = %s is exist \nvalue = %s \n", checkKey, value)
	} else {
		fmt.Println("the key is not exist");
	}
	fmt.Println("--------------------------------------------")

	// how to del the element
	fmt.Printf("del the key => %s", checkKey)
	delete(mapSample2, checkKey)
	fmt.Println(mapSample2)
	fmt.Println("--------------------------------------------")

}
