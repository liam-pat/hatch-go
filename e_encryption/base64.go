package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Printf("bash64 str `%s` encode to `%s` \n", "test_01", base64encodeStr("test_01"))
	fmt.Printf("bash64 str `%s` decode to `%s` \n", base64encodeStr("test_01"), base64DecodeStr(base64encodeStr("test_01")))

}

func base64encodeStr(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func base64DecodeStr(input string) string {
	result, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		_ = fmt.Errorf("decode failed,%v\n", err)
	}

	return string(result)
}
