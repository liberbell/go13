package main

import "fmt"

func main() {
	var any interface{}

	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 3
	fmt.Println(any)
}
