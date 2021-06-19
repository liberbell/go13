package main

import "fmt"

func main() {
	fmt.Println("......ARRAYS")
	arrays()
}

func arrays() {
	nums := [...]int{1, 2, 3}
	incr(nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	for i := range nums {
		nums[i]++
	}
}
