package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
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
		// write
		file, _ := os.Create("/tmp" + "/test.txt")
		defer file.Close()

		f, _ := os.OpenFile("/tmp"+"/test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
		defer f.Close()

		_, _ = f.Write([]byte("hello (by []byte)\n"))
		fileInfo, _ := os.Stat("/tmp" + "/test.txt")
		fmt.Printf("name : %s \nmode : %s,\nsize : %d B \n", fileInfo.Name(), fileInfo.Mode(), fileInfo.Size())
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		absPath, _ := filepath.Abs("/tmp" + "/test.txt")
		f, _ := os.Open(absPath)
		defer f.Close()

		readByte := make([]byte, 128)
		for {
			n, _ := f.Read(readByte)
			fmt.Println(string(readByte[:n]))
			if n < 128 {
				fmt.Println("read end")
				break
			}
		}
		_ = os.Remove(absPath)
	}
	{
		var inputStr string
		if len(os.Args) > 1 {
			inputStr += strings.Join(os.Args[1:], ",")
		}
		fmt.Println("cmd args: ", inputStr)
	}
	{
		sourceFilepath := dir + "/source.txt"

		fmt.Println("dir : ", path.Dir(sourceFilepath))
		fmt.Println("file name : ", path.Base(sourceFilepath))
		fmt.Println("file ext : ", path.Ext(sourceFilepath))
		fmt.Println(strings.Repeat("##", 20))
	}

	{
		srcFile := dir + "/source.txt"
		dstFile := dir + "/copy.txt"

		src, _ := os.Open(srcFile)
		defer src.Close()

		dst, _ := os.Create(dstFile)
		defer dst.Close()

		io.Copy(dst, src)
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		file, _ := os.Open(dir + "/source.txt")
		defer file.Close()

		buf := make([]byte, 1024)
		var content string
		read := bufio.NewReader(file)

		for {
			if n, _ := read.Read(buf); n == 0 {
				break
			}
			content += string(buf[:])
		}
		fmt.Println(content)
		fmt.Println(strings.Repeat("##", 10))
	}
	{
		file, _ := os.Open(dir + "/book.csv")
		defer file.Close()

		books := make([]Book, 1)
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
			line = string(line[:len(line)-1])
			strSlice := strings.Split(line, ";")

			price, _ := strconv.ParseFloat(strSlice[1], 32)
			quantity, _ := strconv.Atoi(strSlice[2])

			book := &Book{strSlice[0], price, quantity}

			if books[0].title == "" {
				books[0] = *book
			} else {
				books = append(books, *book)
			}
		}
		fmt.Println(books)
		fmt.Println(strings.Repeat("##", 10))
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
		fmt.Println(strings.Repeat("##", 10))
	}
	{
		// read file one by one
		file, _ := os.Open(dir + "/book.csv")
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Printf("The input is: %s", str)
		}
		fmt.Println(strings.Repeat("##", 10))
	}
	{
		file := dir + "/data.json"
		buf, _ := os.ReadFile(file)
		fmt.Println("%s", string(buf))
		fmt.Println(strings.Repeat("##", 10))
	}
	{
		// read zip . support bzip2、gzip、lzw 、 zlib。
		file := dir + "/encode.gz"
		filep, _ := os.Open(file)
		defer filep.Close()

		var r *bufio.Reader
		zip, err := gzip.NewReader(filep)
		if err != nil {
			r = bufio.NewReader(filep)
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
