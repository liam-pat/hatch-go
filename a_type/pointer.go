package main

import (
	"fmt"
	"strings"
)

func increase01(nums [5]int) {
	for i := range nums {
		nums[i]++
	}
}
func increase02(nums *[5]int) {
	for i := range nums {
		nums[i]++
	}
}

func up01(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
	list = append(list, "BUG")
}
func up02(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	*list = append(*list, "BUG")
}

func changeMapValue(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// struct
type house struct {
	name  string
	rooms int
}

func addRoom(h house) {
	h.rooms++
}
func addRoomPtr(h *house) {
	h.rooms++ // same: (*h).rooms++
}

func main() {
	{
		nums := [...]int{1, 2, 3, 4, 5}
		fmt.Printf("array %v\t%p\n", nums, &nums) // array [1 2 3 4 5]

		increase01(nums)
		fmt.Printf("nums [5]int  : %v\t\n", nums) // nums [5]int  : [1 2 3 4 5]

		increase02(&nums)
		fmt.Printf("nums *[5]int : %v\t\n", nums) // nums *[5]int : [2 3 4 5 6]

		fmt.Println(strings.Repeat("###", 10))
	}
	{
		dirs := []string{"up", "down", "left", "right"}
		fmt.Printf("list : %p %q\n", &dirs, dirs) // 0x1400000c0c0 ["up" "down" "left" "right"]

		// NOTE, change the existing element value, the original slice will be changed
		// but the append will not affect the original slice
		up01(dirs)
		fmt.Printf("[]string: %p %q\n", &dirs, dirs) // 0x1400000c0c0 ["UP" "DOWN" "LEFT" "RIGHT"]

		up02(&dirs)
		fmt.Printf("*[]string: %p %q\n", &dirs, dirs) // 0x1400000c0c0 ["UP" "DOWN" "LEFT" "RIGHT" "BUG"]

		fmt.Println(strings.Repeat("###", 10))
	}
	{
		confused := map[string]int{"one": 2, "two": 1}
		fmt.Printf("map : %p %v\n", &confused, confused) // 0x1400010c030 map[one:2 two:1]
		changeMapValue(confused)
		fmt.Printf("m map[string]int : %p %v\n", &confused, confused) //0x1400010c030 map[one:1 three:3 two:2]

		fmt.Println(strings.Repeat("###", 10))
	}
	{
		myHouse := house{name: "My House", rooms: 5}
		fmt.Printf("struct: %p %+v\n", &myHouse, myHouse)

		addRoom(myHouse) // copy value
		fmt.Printf("addRoom(myHouse): %p %+v\n", &myHouse, myHouse)

		addRoomPtr(&myHouse)
		fmt.Printf("addRoomPtr(&myHouse): %p %+v\n", &myHouse, myHouse)

		fmt.Println(strings.Repeat("###", 10))
	}
	{
		var counter byte = 100
		P := &counter
		V := *P

		fmt.Printf("counter -> %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       -> %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       -> %-16d address: %-16p\n", V, &V)
		fmt.Println(strings.Repeat("###", 10))

		V = 200
		fmt.Printf("counter -> %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       -> %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       -> %-16d address: %-16p\n", V, &V)
		fmt.Println(strings.Repeat("###", 10))

		V = counter // reset the V to counter's initial value
		counter++
		fmt.Printf("counter -> %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       -> %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       -> %-16d address: %-16p\n", V, &V)
		fmt.Println(strings.Repeat("###", 10))

		*P = 25
		fmt.Printf("counter -> %-16d address: %-16p\n", counter, &counter)
		fmt.Printf("P       -> %-16p address: %-16p *P: %-16d\n", P, &P, *P)
		fmt.Printf("V       -> %-16d address: %-16p\n", V, &V)
		fmt.Println(strings.Repeat("###", 10))
	}
}
