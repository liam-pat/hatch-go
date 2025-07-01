package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {

	var (
		name string
		sex  string
	)

	name, sex = "Alan", "man"
	fmt.Println(name, sex)
	name, sex = sex, name
	fmt.Println(name, sex)
	fmt.Println(int16(1000) + int16(int8(127)))

	var orange int32 = 65 // 65 is A
	fmt.Printf("int: %d 2str: %s\n", orange, string(orange))

	fmt.Println(string([]byte{104, 105}))

	fmt.Println("---------Binary")
	fmt.Printf("%b\n", 2)

	fmt.Printf("%04b = %[1]d\n", 8)
	fmt.Printf("%04b = %[1]d\n", 4)
	fmt.Printf("%04b = %[1]d\n", 2)
	fmt.Printf("%04b = %[1]d\n", 10)
	fmt.Printf("%04b = %[1]d\n", 16)

	fmt.Printf("%08b = %[1]d\n", 16)
	fmt.Printf("%08b = %[1]d\n", 32)
	fmt.Printf("%08b = %[1]d\n", 64)
	fmt.Printf("%08b = %[1]d\n", 128)

	fmt.Println("---------Binary To Decimal")
	i, _ := strconv.ParseInt("00000010", 2, 64)
	fmt.Printf("type = %T, value = %[1]v\n", i)
	i, _ = strconv.ParseInt("00010110", 2, 64)
	fmt.Printf("type = %T, value = %[1]v\n ", i)

	fmt.Println("---------The Max Range")
	fmt.Println("int8  Range :", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 Range :", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 Range :", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 Range :", math.MinInt64, math.MaxInt64)

	fmt.Println("uint8 Range  :", 0, math.MaxUint8)
	fmt.Println("uint16 Range :", 0, math.MaxUint16)
	fmt.Println("uint32 Range :", 0, math.MaxUint32)
	fmt.Println("uint64 Range :", 0, uint64(math.MaxUint64))

	fmt.Println("---------Binary Decimal Map")
	fmt.Printf("%016b %[1]d\n", uint16(65535))

	fmt.Println("---------Parse The Time")
	h, _ := time.ParseDuration("4h30m")
	fmt.Println(h.Hours(), "hours")
	h *= time.Duration(2)
	fmt.Println(h)
}
