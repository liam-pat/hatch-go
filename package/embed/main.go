package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed static/embed_string.log
var fileString string //utf8 string

//go:embed static/embed_byte.log
var fileByte []byte //bin file . eg.images, meta

//go:embed static/*.log
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))
	fmt.Println(strings.Repeat("*#", 20))

	content1, _ := folder.ReadFile("static/embed_string.log")
	fmt.Println(string(content1))
	fmt.Println(strings.Repeat("*#", 20))

	content2, _ := folder.ReadFile("static/embed_byte.log")
	fmt.Println(string(content2))
	fmt.Println(strings.Repeat("*#", 20))
}
