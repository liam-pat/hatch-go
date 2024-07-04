# Purpose

* To study language Go
* Inquiry more conveniently

# Useful CMD

```go
go mod init // init go modules
go get -v github.com/go -ego/gse@v0.60.0-rc4.2 // get the module

/**
  -d -- download but no build
  -u -- force to download/update the package
  -v -- more output info¬
  -t -- get the test package
*/
go get // download the code to scr dir and then go install to gen the bin file to bin dir
go get -u // update the package
go get -u github.com/go -ego/gse // update the module
go get -u github/com/go -ego/gse@v0.60.0-rc4.2 // update the specified version

// replace 
go mod edit -replace github.com/go -ego/gse = /path/to/local/gse
go mod edit -replace github.com/go -ego/gse = github.com/vcaesar/gse

// additional
go mod download // download the independence 
go mod tidy   // update the independence
go mod vendor // save 2 vendor dir
go mod graph  // output the independent graph
go mod verify // verify the modules 

// build & install
go build   // main package -> will gen the file to current dir, common package -> no file will be gen
go install // will gen .a file adn the move to $GOPATH/pkg or $GOPATH/bin
go clean  // remove the building cache files
go format // format the code style

```

# Test Case

```
# run the test case
bash > go test . 

# inspect the coverage
bash > go test -coverprofile=c.out 
bash > go tool cover -html=c.out -o coverage.html 

# inspect cpu resource
bash > go test -bench . -cpuprofile cpu.out 
bash > go tool pprof cpu.out
```




