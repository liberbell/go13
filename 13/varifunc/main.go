package main

import "fmt"

func main() {
	local := 10
	show(local)

	// incrWrong(local)
	fmt.Printf("main -> local = %d\n", local)

	incr(local)
	show(local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incr(n int) int {
	n++
	return n
}
