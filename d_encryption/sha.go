package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "test_01"
	hash := sha1.New()
	hash256 := sha256.New()
	hash512 := sha512.New()

	hash.Write([]byte(str))
	hash256.Write([]byte(str))
	hash512.Write([]byte(str))

	encodeStr1 := hex.EncodeToString(hash.Sum(nil))
	encodeStr2 := hex.EncodeToString(hash256.Sum(nil))
	encodeStr3 := hex.EncodeToString(hash512.Sum(nil))

	fmt.Printf("%-10s `%s`=> `%s` \n", "sha1", "test_01", encodeStr1)
	fmt.Printf("%-10s `%s`=> `%s` \n", "sha256", "test_01", encodeStr2)
	fmt.Printf("%-10s `%s`=> `%s` \n", "sha512", "test_01", encodeStr3)
}
