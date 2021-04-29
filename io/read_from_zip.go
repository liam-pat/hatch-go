package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// support bzip2、flate、gzip、lzw 、 zlib。
	fName := "./io/file/encode.gz"
	var r *bufio.Reader

	file, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	defer file.Close()

	zip, err := gzip.NewReader(file)
	if err != nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(zip)
	}

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println("finish read")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
