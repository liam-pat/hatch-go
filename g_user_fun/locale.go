package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	var lang string // from the url query param or the frontend header `Accept-Language`
	if len(os.Args) != 2 {
		lang = "zh-CN"
	} else {
		lang = os.Args[1]
	}
	locales := make(map[string]map[string]string)

	en := make(map[string]string, 10)
	en["name"] = "liam"
	en["sex"] = "male"
	en["time_zone"] = "America/Chicago"

	location, _ := time.LoadLocation("America/Chicago")
	en["time"] = time.Now().In(location).Format(time.RFC3339)

	cn := make(map[string]string, 10)
	cn["name"] = "李心艾"
	cn["sex"] = "男"
	cn["time_zone"] = "Asia/Shanghai"
	location, _ = time.LoadLocation("Asia/Shanghai")
	cn["time"] = time.Now().In(location).Format(time.RFC3339)

	locales["en-US"] = en
	locales["zh-CN"] = cn

	if v, ok := locales[lang]; ok {
		data, _ := json.Marshal(v)
		fmt.Println(string(data))
		return
	} else {
		log.Fatal(fmt.Printf("cannot find the locale the %s", lang))
	}
}
