package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//  request.URL.Path return /file/xxxxx
	path := request.URL.Path[len("/file/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
