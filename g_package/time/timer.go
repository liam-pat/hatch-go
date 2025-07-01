package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	{
		timer := time.NewTimer(time.Second * 1)
		ch := time.After(time.Second * 3)

		go func() {
			<-timer.C
			fmt.Println("1s later")
		}()
		go func() {
			<-ch
			fmt.Println("3s later")
		}()
		time.Sleep(time.Second * 5)

		fmt.Println(strings.Repeat("###", 30))
	}
	{
		timer := time.NewTimer(2 * time.Second)
		fmt.Println("start time: ", time.Now().Format("2006-01-02 15:04:05"))

		go func() {
			times := 1
			for {
				<-timer.C
				fmt.Printf("--reset time %d, %s will reset 2s\n", times, time.Now().Format("2006-01-02 15:04:05"))
				timer.Reset(2 * time.Second)
				if (times) > 3 {
					fmt.Println("reset time > 3, stop")
					timer.Stop()
				}
				times++
			}
		}()

		time.Sleep(10 * time.Second)
		fmt.Println("end time: ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		go func() {
			for {
				<-ticker.C
				fmt.Println("Ticker accept at :", time.Now().Format("2006-01-02 15:04:05"))
			}
		}()
		time.Sleep(3 * time.Second)
		fmt.Println(strings.Repeat("###", 30))
	}
	{
		done := make(chan int)
		go func() {
			time.Sleep(2 * time.Second)
			<-done
		}()
		select {
		case done <- 1:
			fmt.Println("sth need to do...")
		case <-time.After(2 * time.Second):
			close(done)
			fmt.Println("timeout...")
		}
		fmt.Println(strings.Repeat("###", 30))
	}
}
