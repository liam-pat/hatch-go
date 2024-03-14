# Purpose
> Get more familiar with using GoLand.

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




