package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("plz, input a dir name")
	}
	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	var filename []byte
	for _, file := range files {
		name := file.Name()
		filename = append(filename, name...)
		filename = append(filename, '\n')
		fmt.Println(name)
	}
	err = os.WriteFile(fmt.Sprintf("../tmp/tmp_%d.txt", time.Now().Unix()), filename, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
