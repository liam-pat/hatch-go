# go cmd

# go cmd

```
# bulid . exec file will gen in current dir , if wanna gen to $GOPATH/bin , plz exec go install
# different arch to build (go version >= 1.5) >>>> GO ENABLED=O GOOS=linux GOARCH=amd64 go build xxx.go
bash > go build
# run
bash > go run 
# format code
bash > go fmt
# install
bash > go install 

# test
bash > go test
bash > go test -v -cover
bash > go test -v -run Test_hello test/hello_test.go  // only exec the Test_hello func

# remove the build files 
bash > go clean

# only updates current mod 
bash > go get 
# only updates dependent mods
bash > go get -u
# updates current mod and dependent mods
bash > go get -u -t
# if go get cannot download 
bash > mkdir -p golang.org/x/ && cd  golang.org/x/
bash > git clone https://github.com/golang/sync.git
####
# 3 types to get 
bash > go get golang.org/x/text@latest  // latest version ,if it is tag, will get the tag 
bash > go get golang.org/x/text@master  // master branch latest commit id 
bash > go get golang.org/x/text@v0.32   // get the tag 
bash > go get golang.org/x/text@342b2e	// get the commit id 
####

# clear up the cache 
bash > go clean -modcache

```

## go mod cmd

```
# init the go mod , create the go mod file 
bash > go mod init

# download all go mods
bash > go mod download 

# tidy the go mods
bash > go mod tidy

# view the dependence
bash > go mod graph

# edit the go mod file
bash > go mod edit 

# export all mods to vendor dir
bash > go mod vendor

# verify the go mod changes
bash > go mod verify

# why to get
bash > go mod why
```