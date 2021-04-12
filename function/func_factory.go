package main

import (
	"fmt"
	"strings"
)

func makeAddSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasPrefix(fileName, suffix) {
			return fileName
		}
		return fileName + suffix
	}
}

func main() {
	addPngSuffixFun := makeAddSuffix(".png")
	addJpgSuffixFun := makeAddSuffix(".jpg")

	a := addPngSuffixFun("test")
	b := addJpgSuffixFun("test")

	fmt.Println(a)
	fmt.Println(b)
}
