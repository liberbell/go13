package main

import "fmt"

func main() {
	m, n := 5, 2
	res := multiply(m, n)
	fmt.Println(res)
}

func multiply(a, b int) int {
	r := a * b
	return r
}
