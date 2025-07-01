package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type BookStore struct {
	XMLName     xml.Name `xml:"books"`
	Version     string   `xml:"version,attr"`
	Store       []book   `xml:"book"`
	Description string   `xml:",innerxml"`
}

type book struct {
	XMLName   xml.Name `xml:"book"`
	BookName  string   `xml:"bookName"`
	BookPrice string   `xml:"bookPrice"`
}

func main() {
	{
		xmlStr := `
<?xml version="1.0" encoding="utf-8"?>
<books version="1">
	<book>
		<bookName>math</bookName>
		<bookPrice>120</bookPrice>
	</book>
	<book>
		<bookName>eg</bookName>
		<bookPrice>75</bookPrice>
	</book>
</books>
`
		v := BookStore{}

		if err := xml.Unmarshal([]byte(xmlStr), &v); err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		fmt.Println("struct: ", v)
		fmt.Println(strings.Repeat("##", 20))
	}

	{
		bs := &BookStore{Version: "1"}

		bs.Store = append(bs.Store, book{BookName: "1st", BookPrice: "120"})
		bs.Store = append(bs.Store, book{BookName: "2nd", BookPrice: "75"})

		if output, err := xml.MarshalIndent(bs, "  ", "    "); err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			os.Stdout.Write([]byte(xml.Header))
			os.Stdout.Write(output)
		}
	}
}
