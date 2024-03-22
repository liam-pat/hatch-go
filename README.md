# Purpose

> Get more familiar with using GoLand.

# GoLang execution order

* `main.go` to start and then generate the lookup structure tree
* initialize the tree step by step
* In the initialization of a single package, constants are initialized first, followed by global variables, finally, the
  package's init function is executed.

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




