package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DATA = make(map[string]myElement)

type myElement struct {
	Id   string
	Name string
}

func ADD(key string, value myElement) bool {
	if key == "" {
		return false
	}
	if LOOKUP(key) == nil {
		DATA[key] = value

		return true
	}

	return false
}

func DELETE(key string) bool {
	if key == "" {
		return false
	}
	if LOOKUP(key) != nil {
		delete(DATA, key)

		return true
	}

	return false
}

func LOOKUP(key string) *myElement {

	_, ok := DATA[key]

	if ok {
		element := DATA[key]

		return &element
	}

	return nil
}

func UPDATE(key string, newValue myElement) bool {
	if key == "" {
		return false
	}
	DATA[key] = newValue

	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key ==== %s , value ==== %v\n", k, v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		strKeys := strings.Fields(text)

		switch len(strKeys) {
		case 0:
			continue
		case 1:
			strKeys = append(strKeys, "default-key")
			strKeys = append(strKeys, "default-id")
			strKeys = append(strKeys, "default-name")
		case 2:
			strKeys = append(strKeys, "default-id")
			strKeys = append(strKeys, "default-name")
		case 3:
			strKeys = append(strKeys, "default-name")
		}
		switch strKeys[0] {
		case "PRINT":
			PRINT()
		case "EXIT":
			return
		case "ADD":
			ele := myElement{strKeys[2], strKeys[3]}
			if !ADD(strKeys[1], ele) {
				fmt.Println("Add a element failed")
			}
		case "DELETE":
			if !DELETE(strKeys[1]) {
				fmt.Println("Delete a element failed")
			}
		case "LOOKUP":
			if LOOKUP(strKeys[1]) == nil {
				fmt.Println("Look up a element failed")
			}
		case "UPDATE":
			ele := myElement{strKeys[2], strKeys[3]}
			if !UPDATE(strKeys[1], ele) {
				fmt.Println("Update a element failed")
			}
		default:
			fmt.Println("CMD Support : PRINT,EXIT,ADD,DELETE,LOOKUP,UPDATE")
		}
	}
}
