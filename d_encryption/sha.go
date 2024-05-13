package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Printf("sha1   encode value `%s`=> `%s` \n", "test_01", shaToStr("test_01"))
	fmt.Printf("sha256 encode value `%s`=> `%s` \n", "test_01", sha256ToStr("test_01"))
	fmt.Printf("sha512 encode value `%s`=> `%s` \n", "test_01", sha512ToStr("test_01"))

}

func shaToStr(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}

func sha256ToStr(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}

func sha512ToStr(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}
