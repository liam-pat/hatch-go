package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	{
		hash := md5.New()

		hash.Write([]byte("test_01"))

		md5Str := hex.EncodeToString(hash.Sum(nil))

		// second md5
		str := md5.Sum([]byte(md5Str))

		fmt.Printf("encode = %xs\n", md5Str)
		fmt.Printf("2nd encode = %x\n", str)
	}
}
