package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

/*
*

	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
*/
const dir = "./file"

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	{
		// parse the cmd args
		var inputStr string
		if len(os.Args) > 1 {
			inputStr += strings.Join(os.Args[1:], " ")
		}
		fmt.Println(inputStr)
	}
	{
		//  path package

		sourceFilepath := dir + "/source.txt"

		fmt.Println("dir : ", path.Dir(sourceFilepath))
		fmt.Println("file name : ", path.Base(sourceFilepath))
		fmt.Println("file ext : ", path.Ext(sourceFilepath))
	}

	{
		// copy the file
		sourceFilepath := dir + "/source.txt"
		dstFilepath := dir + "/copy.txt"

		src, _ := os.Open(sourceFilepath)
		defer src.Close()

		dst, _ := os.Create(dstFilepath)
		defer dst.Close()

		io.Copy(dst, src)
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read to buffer
		file, _ := os.Open(dir + "/data.json")

		buf := make([]byte, 1024) // 1kb a times
		var content string

		inputReader := bufio.NewReader(file)

		for {
			n, _ := inputReader.Read(buf)
			if n == 0 {
				break
			}
			content += string(buf[:n])
		}

		fmt.Println(content)
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read csv
		books := make([]Book, 1)
		file, _ := os.Open(dir + "/book.csv")
		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
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
		}
		fmt.Println(books)
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read col
		file, _ := os.Open(dir + "/col.txt")
		defer file.Close()

		var col1, col2, col3 []string
		for {
			var v1, v2, v3 string
			_, err := fmt.Fscanln(file, &v1, &v2, &v3)
			if err != nil {
				break
			}
			col1 = append(col1, v1)
			col2 = append(col2, v2)
			col3 = append(col3, v3)
		}

		fmt.Println(col1, col2, col3)
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read file one by one
		file, _ := os.Open(dir + "/book.csv")
		defer file.Close()

		inputReader := bufio.NewReader(file)
		for {
			readString, readerError := inputReader.ReadString('\n')
			if readerError == io.EOF {
				break
			}
			fmt.Printf("The input is: %s\n", readString)
		}
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read content one time
		inputFile := dir + "/data.json"

		buf, _ := ioutil.ReadFile(inputFile)
		fmt.Printf("%s\n", string(buf))
		fmt.Println(strings.Repeat("*#", 10))
	}
	{
		// read zip . support bzip2、gzip、lzw 、 zlib。
		fName := dir + "/encode.gz"
		var r *bufio.Reader

		file, _ := os.Open(fName)
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
				break
			}
			fmt.Println(line)
		}
	}

}
