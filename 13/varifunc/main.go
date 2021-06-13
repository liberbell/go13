package main

import "fmt"

func main() {
	local := 10
	show(local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}
