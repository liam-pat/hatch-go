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
			fmt.Printf("error: %v", err)
			return
		}

		fmt.Println("decode result : ", v)
		fmt.Println(strings.Repeat("##", 20))
	}

	{
		bs := &BookStore{Version: "1"}
		bs.Store = append(bs.Store, book{BookName: "test01", BookPrice: "120"})
		bs.Store = append(bs.Store, book{BookName: "test03", BookPrice: "75"})

		if output, err := xml.MarshalIndent(bs, "  ", "    "); err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			os.Stdout.Write([]byte(xml.Header))
			os.Stdout.Write(output)
		}
	}
}
