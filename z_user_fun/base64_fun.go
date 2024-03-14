package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	base64StrEncode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("string `%s` base64 encode 2 string = %s\n", data, base64StrEncode)
	fmt.Println(strings.Repeat("----", 20))

	sDec, _ := base64.StdEncoding.DecodeString(base64StrEncode)
	fmt.Printf("base64 string `%s` decode 2 string = %s\n", base64StrEncode, string(sDec))
	fmt.Println(strings.Repeat("----", 20))

	base64UrlStrEncode := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Printf("string `%s` base64URL encode 2 string = %s\n", data, base64UrlStrEncode)
	fmt.Println(strings.Repeat("----", 20))

	uDec, _ := base64.URLEncoding.DecodeString(base64UrlStrEncode)
	fmt.Printf("base64URL string `%s` decode 2 string = %s\n", base64StrEncode, string(uDec))
	fmt.Println(strings.Repeat("----", 20))
}
