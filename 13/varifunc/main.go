package main

import "fmt"

func main() {
	local := 10
	show()
}

func show() int {
	fmt.Printf("show -> n = %d\n", local)
}
