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
