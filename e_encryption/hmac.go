package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

/*
*
Hash-based Message Authentication Code
*/
func main() {

	fmt.Printf("hamc md5  encode value `%s`=> `%s` \n", "test_01", hmacWithMD5ToStr("test_01", "secret_key"))
	fmt.Printf("hamc sha1 encode value `%s`=> `%s` \n", "test_01", hmacWithSHA1ToStr("test_01", "secret_key"))
	fmt.Printf("hamc sha256 encode value `%s`=> `%s` \n", "test_01", hmacWithSHA256ToStr("test_01", "secret_key"))
	fmt.Printf("hamc sha512 encode value `%s`=> `%s` \n", "test_01", hmacWithSHA512ToStr("test_01", "secret_key"))

}

func hmacWithMD5ToStr(input, key string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum([]byte("")))
}

func hmacWithSHA1ToStr(input, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func hmacWithSHA256ToStr(input, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func hmacWithSHA512ToStr(input, key string) string {
	hash := hmac.New(sha512.New, []byte(key))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte("")))
}
