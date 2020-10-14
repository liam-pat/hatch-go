package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//最后把缓存区的内容输入到文件
	defer writer.Flush()

	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	writeFile("./handle_file/write_file/test.txt")
}
