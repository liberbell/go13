package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("......ARRAYS")
	arrays()

	fmt.Println("......SLICES")
	slices()

	fmt.Println("......MAPS")
	maps()

	fmt.Println("......STRUCTS")
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: " My House", rooms: 5}
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	// fmt.Println(dirs)
	fmt.Printf("slices list    : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list    : %p %q\n", &dirs, dirs)
}

func upPtr(list *[]string) {
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	*list = append(*list, "HEISEN BUG")
	fmt.Printf("upPtr list     : %p %q\n", &list, list)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, "HEISEN BUG")
	fmt.Printf("up_list        : %p %q\n", &list, list)
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("arrays nums    : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)

}

func incr(nums [3]int) {
	fmt.Printf("incr nums    : %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums    : %p\n", &nums)
	for i := range nums {
		(*nums)[i]++
	}
}
