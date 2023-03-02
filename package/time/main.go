package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		nowTime := time.Now()
		fmt.Println(nowTime)
		fmt.Println(time.Time{}.Location())

		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		customTime := time.Date(2008, 7, 15, 13, 30, 0, 0, time.Local)
		fmt.Println(customTime)
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		nowTime1 := time.Now()
		stringTime := nowTime1.Format("2006-1-2- 15:04:05")
		fmt.Println(stringTime)
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// parse time
		strTime := "2019-01-01 15:03:01"
		objTime, _ := time.Parse("2006-01-02 15:04:05", strTime)
		fmt.Println(objTime)
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		nowTime2 := time.Now()
		year, month, day := nowTime2.Date()
		fmt.Printf("%d-%s-%d\n", year, month, day)

		hour, min, sec := nowTime2.Clock()
		fmt.Printf("%d:%d:%d\n", hour, min, sec)
		fmt.Printf("%d-%s-%d\n", nowTime2.Year(), nowTime2.Month(), nowTime2.Hour())
		fmt.Println("already consume : ", nowTime2.YearDay())
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		fmt.Println("unix timestamp : ", time.Now().Unix())
		fmt.Println("unix nano timestamp : ", time.Now().UnixNano())

		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		fmt.Println("10s after : ", time.Now().Add(time.Second*10)) // 10s after
		fmt.Println("1 year later : ", time.Now().AddDate(1, 0, 0)) // 1 year later
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// 整點
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
		fmt.Println(t.Truncate(1 * time.Hour))
		fmt.Println(t.Round(1 * time.Hour))

		fmt.Println(t.Truncate(1 * time.Minute))
		fmt.Println(t.Round(1 * time.Minute))
		fmt.Println(strings.Repeat("#*", 10))
	}
	{
		// 定時器 01
		timer := time.NewTimer(time.Second * 3)
		ch := time.After(time.Second * 5)

		go func() {
			<-timer.C
			fmt.Println("3s later")
		}()

		go func() {
			<-ch
			fmt.Println("5s later")
		}()
		time.Sleep(time.Second * 10)
		fmt.Println(strings.Repeat("#*", 20))
	}
	{
		// 定時器 02
		timer := time.NewTimer(5 * time.Second)
		fmt.Println("start time：", time.Now().Format("2006-01-02 15:04:05"))
		go func() {
			times := 0
			for {
				<-timer.C
				fmt.Println("--times : ", times, time.Now().Format("2006-01-02 15:04:05"))
				timer.Reset(2 * time.Second)
				fmt.Println("--reset 2 seconds")
				times++
				if (times) > 3 {
					fmt.Println("times > 3 , stop timer")
					timer.Stop()
				}
			}
		}()
		time.Sleep(20 * time.Second)
		fmt.Println("end time：", time.Now().Format("2006-01-02 15:04:05"))
	}
	{
		// main process end will stop the ticker
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		go func() {
			for {
				<-ticker.C
				fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
			}
		}()
		time.Sleep(30 * time.Second)
	}

}
