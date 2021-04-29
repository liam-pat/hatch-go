package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	hash := sha1.New()
	io.WriteString(hash, "test")

	var b []byte
	fmt.Printf("Result: %x\n", hash.Sum(b))
	fmt.Printf("Result: %d\n", hash.Sum(b))

	hash.Reset()

	data := []byte("We shall overcome!")
	hash.Write(data)
	fmt.Printf("Result: %x\n", hash.Sum(b))
}
