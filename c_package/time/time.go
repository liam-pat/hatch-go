package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		// output timestamp
		nowTime := time.Now()
		fmt.Println("time.Now() : ", nowTime)
		fmt.Println("time location : ", nowTime.Location())
		fmt.Println("Unix : ", nowTime.Unix())
		fmt.Println("UnixMilli : ", nowTime.UnixMilli())
		fmt.Println("UnixNano : ", nowTime.UnixNano())
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		// init date
		fmt.Println("time.Date : ", time.Date(2008, 7, 15, 13, 30, 0, 0, time.Local))
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		// output format
		fmt.Println("time.Now().Format : ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		// parse time
		utcTime, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 15:03:01") // timezone is UTC+0
		localTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 15:03:01", time.Local)
		fmt.Println("parse time via time.Parse ", utcTime)
		fmt.Println("parse time via time.ParseInLocation ", localTime)
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		year, month, day := time.Now().Date()
		fmt.Printf("%d-%s-%d\n", year, month, day)

		hour, minutes, sec := time.Now().Clock()
		fmt.Printf("%d:%d:%d\n", hour, minutes, sec)

		fmt.Printf("%d-%s-%d\n", time.Now().Year(), time.Now().Month(), time.Now().Hour())

		fmt.Println("pass days : ", time.Now().YearDay())
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		fmt.Println("time.Now().Add() + 10s : ", time.Now().Add(time.Second*10)) // 10s after
		fmt.Println("time.Now().AddDate + 1Y : ", time.Now().AddDate(1, 0, 0))   // 1 year later
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
		fmt.Println("t.Truncate(1 * time.Hour) : ", t.Truncate(1*time.Hour))
		fmt.Println("t.Round(1 * time.Hour) : ", t.Round(1*time.Hour))

		fmt.Println("t.Truncate(1 * time.Minute) : ", t.Truncate(1*time.Minute))
		fmt.Println("Round(1 * time.Minute) : ", t.Round(1*time.Minute))

		fmt.Println(strings.Repeat("###", 30))
	}
}
