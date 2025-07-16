package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// go run foo -enable -name=joe a1 a2
	// go run foo -level 8 a1
	// go run
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "null", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("unexpected subcmd")
		os.Exit(1)
	}
	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Printf("subcmd : %s , enable: %t , name: %s\n", os.Args[1], *fooEnable, *fooName)
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Printf("subcmd : %s , level: %d\n", os.Args[1], *barLevel)
	default:
		fmt.Println("unexpected subcmd")
		os.Exit(1)
	}
}
