package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	str := "abc123!?$*&()'-=@~"

	encode64 := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Printf("str `%s` base64 encode = %s\n", str, encode64)
	fmt.Println(strings.Repeat("----", 20))

	sDec, _ := base64.StdEncoding.DecodeString(encode64)
	fmt.Printf("base64 str `%s` decode = %s\n", encode64, string(sDec))
	fmt.Println(strings.Repeat("----", 20))

	urlencode64 := base64.URLEncoding.EncodeToString([]byte(str))
	fmt.Printf("str `%s` base64URL encode = %s\n", str, urlencode64)
	fmt.Println(strings.Repeat("----", 20))

	uDec, _ := base64.URLEncoding.DecodeString(urlencode64)
	fmt.Printf("base64URL str `%s` decode = %s\n", urlencode64, string(uDec))
	fmt.Println(strings.Repeat("----", 20))
}
