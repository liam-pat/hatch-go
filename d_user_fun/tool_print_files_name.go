package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

/*
// generate the xxx.txt to convert many ts fills to 1 mp4 file
// bash > ffmpeg -f concat -safe 0 -i xxx.txt -c copy out.mp4
*/
func main() {
	dirname := "."
	listRes, err := os.ReadDir(dirname)
	if err != nil {
		return
	}
	validList := make([]fs.FileInfo, 0, len(listRes))
	for _, file := range listRes {
		info, err := file.Info()
		if err != nil {
			continue
		}
		if strings.HasPrefix(file.Name(), ".") || !strings.HasSuffix(file.Name(), ".ts") {
			continue
		}
		validList = append(validList, info)
	}
	if len(validList) <= 0 {
		fmt.Println("Cannot find any files")
		return
	}

	exportFilename := "ts2mp4.txt"
	file, err := os.OpenFile(exportFilename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Create the file failed", err)
	}
	write := bufio.NewWriter(file)
	for _, file := range validList {
		lineMsg := fmt.Sprintf("file '%s'\n", file.Name())
		fmt.Printf(lineMsg)
		write.WriteString(lineMsg)
	}
	write.Flush()
}
