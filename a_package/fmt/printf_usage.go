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

	var orange int32 = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)

	fmt.Println(string([]byte{104, 105}))

	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))

	fmt.Println("---------bit")
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	fmt.Println("---------bit")
	fmt.Printf("%02b = %d\n", 0, 0)
	fmt.Printf("%02b = %d\n", 1, 1)
	fmt.Printf("%02b = %d\n", 2, 2)
	fmt.Printf("%02b = %d\n", 3, 3)

	fmt.Println("---------bit")
	fmt.Printf("%08b = %d\n", 1, 1)
	fmt.Printf("%08b = %d\n", 2, 2)
	fmt.Printf("%08b = %d\n", 4, 4)
	fmt.Printf("%08b = %d\n", 8, 8)
	fmt.Printf("%08b = %d\n", 16, 16)
	fmt.Printf("%08b = %d\n", 32, 32)
	fmt.Printf("%08b = %d\n", 64, 64)
	fmt.Printf("%08b = %d\n", 128, 128)

	fmt.Println("---------Binary To Decimal")
	i, _ := strconv.ParseInt("00000010", 2, 64)
	fmt.Println(i)
	i, _ = strconv.ParseInt("00010110", 2, 64)
	fmt.Println(i)

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
	big := uint16(65535)
	small := uint8(big)
	fmt.Printf("%016b %[1]d\n", big)
	fmt.Printf("%016b %[1]d\n", small)

	fmt.Println("---------Parse The Time")
	h, _ := time.ParseDuration("4h30m")
	fmt.Println(h.Hours(), "hours")
	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))
}
