package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

/*
// generate the xxx.txt to convert many ts fills to 1 mp4 file
// bash > ffmpeg -f concat -safe 0 -i xxx.txt -c copy out.mp4
*/
func main() {
	dirname := "."
	list, _ := os.ReadDir(dirname)

	// get all the .ts files
	var tsFiles []string
	for _, file := range list {
		if strings.HasPrefix(file.Name(), ".") || !strings.HasSuffix(file.Name(), ".ts") {
			continue
		}
		tsFiles = append(tsFiles, file.Name())
	}
	if len(tsFiles) <= 0 {
		fmt.Println("Cannot find any .ts files")
		return
	}
	sort.Strings(tsFiles)

	// output the file names to a text file
	file, err := os.OpenFile("/tmp/ts_files.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("Create the file failed", err)
		return
	}
	write := bufio.NewWriter(file)
	for _, file := range tsFiles {
		fmt.Printf(fmt.Sprintf("file '%s'\n", file))
		write.WriteString(fmt.Sprintf("file '%s'\n", file))
	}
	write.Flush()
}
