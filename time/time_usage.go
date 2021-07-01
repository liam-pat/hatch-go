package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	nowTime := time.Now()
	fmt.Println(nowTime)

	// init time
	customTime := time.Date(2008, 7, 15, 13, 30, 0, 0, time.Local)
	fmt.Println(customTime)

	// format time
	nowTime1 := time.Now()
	stringTime := nowTime1.Format("2006-1-2- 15:04:05")
	fmt.Println(stringTime)

	// parse time
	strTime := "2019-01-01 15:03:01"
	objTime, _ := time.Parse("2006-01-02 15:04:05", strTime)
	fmt.Println(objTime)

	// get year month day
	nowTime2 := time.Now()
	year, month, day := nowTime2.Date()
	fmt.Printf("%d-%s-%d\n", year, month, day)

	hour, min, sec := nowTime2.Clock()
	fmt.Printf("%d:%d:%d\n", hour, min, sec)

	fmt.Printf("%d-%s-%d\n", nowTime2.Year(), nowTime2.Month(), nowTime2.Hour())

	// already consume
	fmt.Println("already consume : ", nowTime2.YearDay())

	// get unix timestamp
	fmt.Println("unix timestamp : ", time.Now().Unix())

	//time range
	fmt.Println("10s after : ", time.Now().Add(time.Second*10)) // 10s after
	fmt.Println("1 year later : ", time.Now().AddDate(1, 0, 0)) // 1 year later

}
