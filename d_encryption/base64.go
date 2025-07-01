package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	{
		encodedStr := base64.StdEncoding.EncodeToString([]byte("test_01"))
		DecodeString, _ := base64.StdEncoding.DecodeString(encodedStr)

		fmt.Printf("bash64 str `%s` encode to `%s` \n", "test_01", encodedStr)
		fmt.Printf("bash64 str `%s` decode to `%s` \n", encodedStr, string(DecodeString))
	}
}
