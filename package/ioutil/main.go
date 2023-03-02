package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const tmpDir = "../../tmp"

func listFileTree(dirStr string, times int) {
	fileInfos, err := ioutil.ReadDir(dirStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			if strings.Contains(info.Name(), ".") {
				continue
			}
			for tmp := times; tmp > 0; tmp-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), ">>>")
			listFileTree(dirStr+"/"+info.Name(), times+1)
		} else {
			for tmp := times; tmp > 0; tmp-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

func main() {
	{
		listFileTree("../", 0)
	}
	{
		//read
		ret, _ := ioutil.ReadFile(tmpDir + "/iotil.txt")
		fmt.Println(string(ret))
	}
}
