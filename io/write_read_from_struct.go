package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() (err error) {
	return ioutil.WriteFile(page.Title, page.Body, 0666)
}

func (page *Page) load(title string) (err error) {
	page.Title = title
	page.Body, err = ioutil.ReadFile(page.Title)
	return err
}

func main() {
	//save into file
	page := Page{
		"./io/file/readme.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var newPage Page
	newPage.load("./io/file/readme.md")
	fmt.Println(string(newPage.Body))
}
