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
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Println(dirs)
	fmt.Printf("slices list    : %p %q\n", &dirs, dirs)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, "HEISEN BUG")
	fmt.Printf("up_list      : %p %q\n", &list, list)
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("array nums    : %p\n", &nums)
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
