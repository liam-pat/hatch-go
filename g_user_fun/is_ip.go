package main

import (
	"fmt"
	"regexp"
)

func isIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	str := "192.168.55.1"
	fmt.Printf("`%s` is ip: %t \n", str, isIP(str))
}
