package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	fmt.Printf("way 1 md5 : %s \n", md5toStr("test_01"))
	fmt.Printf("way 2 md5 : %s \n", md5toStr2("test_01"))
}

func md5toStr(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))

}

func md5toStr2(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
