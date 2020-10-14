
# Project

## goLang 环境

* Mac Ox: 在官方下载后一般放在，`~/go/`

> 一共有三个文件
> 
> src : 放源代码（建议把所有项目放在这里，引用包的时候 ide 会自动引入）
> 
> bin : 可执行的程序，所以这个路径需要放到系统的 $PATH
> 
> pkg : 中间生成的代码（一般不需要管）
> 
> go env 可以查看 go 配置和环境
> 
> go run 直接运行代码，不生成中间代码和可执行文件
> 
> go build 生成可执行代码
> 
> go install 生成可执行文件和中间文件

## goLand 的目录结构

### .go in Same Dir(同级目录)

   * 分文件编程必须放在 src 里面
   
   * 设置 GOPATH 环境变量
   
   * 同一个目录，包名必须一样
   
   * go env 在 terminal 可以查看 go 的环境变量
   
   * 同一个目录，调用其他文件的函数直接调用，不需要引用包名

### .go in Diff Dir（不同级目录）

   * 不同目录，包名不一样
   
   * 调用不同包里面的函数，eg. package.MyFunction()
   
   * 调用别的包里面的函数，必须确认别的包里面的方法是允许外部调用的，go 默认规则是小写字母开头的为b private,大写字母为 public
   
   
### terminal run test case

```
#跑当前的目录的测试
go test . 

#查看覆盖率
go test -coverprofile=c.out 
go tool cover -html=c.out -o coverage.html 

#查看 cpu 占用时间
go test -bench . -cpuprofile cpu.out 
go tool pprof cpu.out
>>web
```
    
  ### generate the go DOC
    
```bash

#生成 go Doc
cd godoc

go doc

go doc `YOUR SREUCT`

godoc -http :6060

```




