package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	var b []byte
	md5 := md5.New()

	io.WriteString(md5, "test")
	fmt.Printf("Result: %x\n", md5.Sum(b))
	fmt.Printf("Result: %d\n", md5.Sum(b))

	md5.Reset()

	c := []byte("We shall overcome!")
	md5.Write(c)
	fmt.Printf("Result: %x\n", md5.Sum(c))
	fmt.Printf("Result: %d\n", md5.Sum(c))
}
