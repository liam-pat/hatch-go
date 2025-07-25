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
		log.Fatal("plz input the file name!!!")
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}

	var filename []byte
	for _, file := range files {
		name := file.Name()
		filename = append(filename, name...)
		filename = append(filename, '\n')
		fmt.Println(name)
	}

	file := fmt.Sprintf("/tmp/tmp_%d.txt", time.Now().Unix())
	os.WriteFile(file, filename, 0664)
}
