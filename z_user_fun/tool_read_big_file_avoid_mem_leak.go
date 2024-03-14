package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, slicePool *sync.Pool, start time.Time, end time.Time) {
	//another wait group to process every chunk
	var waitGroup sync.WaitGroup
	logs := stringPool.Get().(string)
	logs = string(chunk)
	linesPool.Put(chunk) //put back the chunk in pool

	//split the string by "\n", so that we have slice of logs
	logsSlice := strings.Split(logs, "\n")
	stringPool.Put(logs) //put back the string pool

	chunkSize := 100 //process the bunch of 100 logs in thread
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}
	length := len(logsSlice)
	//traverse the chunk
	for i := 0; i < length; i += chunkSize {
		waitGroup.Add(1)
		go func(s int, e int) {
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				logParts := strings.SplitN(text, ",", 2)
				logCreationTimeString := logParts[0]
				logCreationTime, err := time.Parse("2006-01-  02T15:04:05.0000Z", logCreationTimeString)
				if err != nil {
					fmt.Printf("\n Could not able to parse the time :%s       for log : %v", logCreationTimeString, text)
					return
				}
				if logCreationTime.After(start) && logCreationTime.Before(end) {

					fmt.Println(text)
				}
			}
			waitGroup.Done()

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}
	waitGroup.Wait()
	logsSlice = nil
}

func main() {
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}
	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}
	slicePool := sync.Pool{New: func() interface{} {
		lines := make([]string, 100)
		return lines
	}}
	filename := "../"
	file, err := os.Open(filename)
	reader := bufio.NewReader(file)
	if err != nil {
		fmt.Printf("open file error : %s \n", err)
		return
	}

	var waitGroup sync.WaitGroup //wait group to keep track off all threads
	for {
		buf := linesPool.Get().([]byte)
		n, err1 := reader.Read(buf)
		buf = buf[:n]
		// if error or ending
		if n == 0 {
			if err1 == io.EOF {
				break
			}
			if err1 != nil {
				fmt.Println(err1)
				break
			}
			return
		}
		nextUntilNewline, err2 := reader.ReadBytes('\n') //read entire line
		if err2 != io.EOF {
			buf = append(buf, nextUntilNewline...)
		}

		waitGroup.Add(1)
		go func() {
			//start -> log start time, end -> log end time
			during, _ := time.ParseDuration("20s")
			start := time.Now()
			end := start.Add(during)

			ProcessChunk(buf, &linesPool, &stringPool, &slicePool, start, end)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}
