package main

import "fmt"

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("main -> n = %d\n", n)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}
