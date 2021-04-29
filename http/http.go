package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.apple.com", nil)

	request.Header.Add("User-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect: ", req)
			return nil
		},
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseString, err := httputil.DumpResponse(response, true)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", responseString)
}
