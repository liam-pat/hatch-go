package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	{
		args := os.Args[1:]
		fmt.Println("all argus :", args)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		wordPointer := flag.String("word", "default_foo", "type : string")
		numbPointer := flag.Int("number", 0, "type : number")
		forkPointer := flag.Bool("fork", false, "type : boolean")

		var name string
		flag.StringVar(&name, "name", "default_name", "type : string")

		flag.Parse()

		fmt.Println("word:", *wordPointer)
		fmt.Println("number:", *numbPointer)
		fmt.Println("fork:", *forkPointer)
		fmt.Println("name:", name)

		fmt.Println(strings.Repeat("##", 20))
	}
}
