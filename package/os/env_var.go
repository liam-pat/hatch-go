package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("name", "package")
	fmt.Println("name:", os.Getenv("name"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
