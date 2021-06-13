package main

import (
	"fmt"
	"strconv"
)

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("main -> local = %d\n", local)

	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	_, err := incrByStr(local, "TWO")
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

func incrByStr(n, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}
