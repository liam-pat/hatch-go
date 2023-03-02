package main

import (
	"fmt"
	"strings"
)

// array function
func array() {
	nums := [...]int{1, 2, 3, 4, 5}

	fmt.Printf("original array %v\t%p\n", nums, &nums)
	fmt.Println(strings.Repeat("*#", 20))

	inCreate(nums)
	fmt.Printf("array handle  : %v\t\n", nums)
	fmt.Println(strings.Repeat("*#", 20))

	inCreateByPoi(&nums)
	fmt.Printf("array handle by pointer : %v\t\n", nums)
	fmt.Println(strings.Repeat("*#", 20))
}
func inCreate(nums [5]int) {
	fmt.Printf("internal address  : %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}
func inCreateByPoi(nums *[5]int) {
	fmt.Printf("internal address  : %p\n", &*nums)
	for i := range nums {
		nums[i]++
	}
}

// slice
func slices() {
	dirs := []string{"up", "down", "left", "right"}
	fmt.Printf("original list      : %p %q\n", &dirs, dirs)
	fmt.Println(strings.Repeat("*#", 30))

	up(dirs)
	fmt.Printf("outside slice list : %p %q\n", &dirs, dirs)
	fmt.Println(strings.Repeat("*#", 30))

	upByPtr(&dirs)
	fmt.Printf("outside point list : %p %q\n", &dirs, dirs)
	fmt.Println(strings.Repeat("*#", 30))
}
func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
	list = append(list, "BUG")
	fmt.Printf("internal slice     : %p %q\n", &list, list)
}
func upByPtr(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	*list = append(*list, "BUG")
	fmt.Printf("internal slice point: %p %q\n", &*list, list)
}

// map
func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fmt.Printf("outside : %p %v\n", &confused, confused)
	fix(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Printf("internal : %p %v\n", &m, m)
}

// struct
type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}
	fmt.Printf("orignal struct       : %p %+v\n", &myHouse, myHouse)
	fmt.Println(strings.Repeat("*#", 10))
	addRoom(myHouse)
	fmt.Printf("outside struct       : %p %+v\n", &myHouse, myHouse)
	fmt.Println(strings.Repeat("*#", 10))
	addRoomPtr(&myHouse)
	fmt.Printf("outside struct pst   : %p %+v\n", &myHouse, myHouse)
}
func addRoom(h house) {
	fmt.Printf("internal struct      : %p %+v\n", &h, h)
	h.rooms++
}
func addRoomPtr(h *house) {
	h.rooms++ // same: (*h).rooms++
	fmt.Printf("internal struct prt  : %p %+v\n", h, h)
}

func main() {
	{
		var counter byte = 100
		P := &counter
		V := *P

		fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

		V = 200
		fmt.Println(strings.Repeat("*#", 10))
		fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

		fmt.Println(strings.Repeat("*#", 10))
		V = counter // reset the V to counter's initial value
		counter++
		fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

		fmt.Println(strings.Repeat("*#", 10))
		*P = 25
		fmt.Println()
		fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       : %-16d address: %-16p\n", V, &V)
	}
	{
		fmt.Println(strings.Repeat("*#", 20))
		array()
		fmt.Println(strings.Repeat("*#", 10))
		// slice 會clone 一個 slice 進入函數，函數內修改 slice 的值， 原來的 slice 也會改變
		// 但是新的 slice 擴容或者新增 element 並不會影響原來的 slice
		slices()
		fmt.Println(strings.Repeat("*#", 10))
		// 地址傳遞
		maps()
		fmt.Println(strings.Repeat("*#", 10))
		// same as array
		structs()
		fmt.Println(strings.Repeat("*#", 10))
	}

}
