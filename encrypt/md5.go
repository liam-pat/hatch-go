package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	hash := md5.New()
	var b []byte

	io.WriteString(hash, "test")
	fmt.Printf("Result: %x\n", hash.Sum(b))
	fmt.Printf("Result: %d\n", hash.Sum(b))

	hash.Reset()

	c := []byte("We shall overcome!")
	hash.Write(c)
	fmt.Printf("Result: %x\n", hash.Sum(c))
	fmt.Printf("Result: %d\n", hash.Sum(c))
}
