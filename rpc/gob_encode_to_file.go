package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type VCardDemo struct {
	FirstName string
	LastName  string
	Remark    string
}

func main() {
	file, _ := os.Open("rpc/file/vcard.gob")
	defer file.Close()

	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)

	var vc VCardDemo
	err := dec.Decode(&vc)

	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
