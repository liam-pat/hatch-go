// read_csv.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	books := make([]Book, 1)

	file, err := os.Open("./io/file/book.csv")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		readErr := err
		// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
		line = string(line[:len(line)-1])

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, _ = strconv.ParseFloat(strSl[1], 32)
		book.quantity, _ = strconv.Atoi(strSl[2])

		if books[0].title == "" {
			books[0] = *book
		} else {
			books = append(books, *book)
		}
		// read to the end
		if readErr == io.EOF {
			break
		}
	}
	// print the books
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range books {
		fmt.Println(bk)
	}
}
