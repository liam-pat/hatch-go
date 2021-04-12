package main

import (
	"fmt"
	"time"
)

// three ways to delay
func main() {
	timer01()
	timer02()
	timer03()
}

func timer01() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func timer02() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}

func timer03() {
	<-time.After(2 * time.Second)
	fmt.Println("时间到")
}
