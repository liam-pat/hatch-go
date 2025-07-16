package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type user struct {
	name string
}

func main() {
	{
		u := user{"tang"}
		fmt.Printf("% + v\n", u)           //格式化输出结构
		fmt.Printf("%%#v = %#v\n", u)      //输出值的 Go 语言表示方法
		fmt.Printf("%%p = %p\n", &u)       //pointer
		fmt.Printf("%%T = %T\n", u)        //输出值的类型的 Go 语言表示
		fmt.Printf("%%t = %t\n", true)     //输出值的 true 或 false
		fmt.Printf("%%q = %q\n", "dedede") // "" quoted string
		fmt.Printf("%%c = %c\n", 67)       //char
		fmt.Printf("%%U = %U\n", 123)      //Unicode

		fmt.Printf("%%b = %b\n", 1023) //二进制表示
		fmt.Printf("%%d = %d\n", 1023) //十进制表示
		fmt.Printf("%%o = %o\n", 1023) //八进制表示
		fmt.Printf("%%x = %x\n", 1023) //十六进制表示，用a-f表示
		fmt.Printf("%%X = %X\n", 1023) //十六进制表示，用A-F表示

		fmt.Printf("%%b = %b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
		fmt.Printf("%%e = %e\n", 12.345)   //科学计数法，e表示
		fmt.Printf("%%E = %E\n", 12.34455) //科学计数法，E表示
		fmt.Printf("%%f = %f\n", 12.3456)  //有小数部分，无指数部分
		fmt.Printf("%%g = %g\n", 12.3456)  //根据实际情况采用%e或%f输出
		fmt.Printf("%%G = %G\n", 12.3456)  //根据实际情况采用%E或%f输出

		fmt.Printf("%%s = %s\n", []byte("test"))  //直接输出字符串或者[]byte
		fmt.Printf("%%x = % x\n", []byte("abcz")) //每个字节用两字节十六进制表示，a-f表示
		fmt.Printf("%%X = % X\n", []byte("abcz")) //每个字节用两字节十六进制表示，A-F表示
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		name, sex := "Alan", "man"
		fmt.Println(name, sex)
		name, sex = sex, name
		fmt.Println(name, sex)

		fmt.Printf("string([]byte{104, 105} = %s\n", string([]byte{104, 105}))

		fmt.Printf("%08b = %[1]d\n", 8)
		fmt.Printf("%08b = %[1]d\n", 15)
		fmt.Printf("%08b = %[1]d\n", 16)
		fmt.Printf("%08b = %[1]d\n", 128)

		i, _ := strconv.ParseInt("00000010", 2, 64)
		fmt.Printf("type = %T, value = %[1]v\n", i)
		i, _ = strconv.ParseInt("00010110", 2, 64)
		fmt.Printf("type = %T, value = %[1]v\n ", i)

		fmt.Println("int8  Range :", math.MinInt8, math.MaxInt8)
		fmt.Println("int16 Range :", math.MinInt16, math.MaxInt16)
		fmt.Println("int32 Range :", math.MinInt32, math.MaxInt32)
		fmt.Println("int64 Range :", math.MinInt64, math.MaxInt64)

		fmt.Println("uint8 Range  :", 0, math.MaxUint8)
		fmt.Println("uint16 Range :", 0, math.MaxUint16)
		fmt.Println("uint32 Range :", 0, math.MaxUint32)
		fmt.Println("uint64 Range :", 0, uint64(math.MaxUint64))

		fmt.Printf("%016b %[1]d\n", uint16(65535))

		h, _ := time.ParseDuration("4h30m")
		fmt.Println(h.Hours(), "hours")
		h *= time.Duration(2)
		fmt.Println(h)
	}
}
