package main

import "fmt"

func main() {
	// var n int
	// n = n + 1
	// fmt.Println(n)

	// n += 1
	// fmt.Println(n)

	// n++
	// fmt.Println(n)

	n := 10
	n = n - 1
	fmt.Println(n)

	n -= 1
	fmt.Println(n)

	n--
	fmt.Println(n)

	var counter int
	counter++
	counter++
	counter++
	counter--

	fmt.Printf("There are %d line(s) in the field.\n", counter)
}
