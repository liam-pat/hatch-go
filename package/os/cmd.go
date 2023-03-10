package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	{
		argsWithoutProg := os.Args[1:]
		fmt.Println("all argument :", argsWithoutProg)
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		wordPointer := flag.String("word", "foo", "type : string")
		numbPointer := flag.Int("numb", 42, "type : number")
		forkPointer := flag.Bool("fork", false, "type : boolean")
		var name string
		flag.StringVar(&name, "name", "default", "type : string")

		flag.Parse()

		fmt.Println("word:", *wordPointer)
		fmt.Println("numb:", *numbPointer)
		fmt.Println("fork:", *forkPointer)
		fmt.Println("name:", name)

		fmt.Println(strings.Repeat("#*", 20))
	}
}
