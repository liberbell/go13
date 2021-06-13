package main

import "fmt"

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("main -> local = %d\n", local)

	local = incr(local)
	show(local)

	local = incrBy(local, 5)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n * factor
}
