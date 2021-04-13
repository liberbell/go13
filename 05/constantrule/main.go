package main

import "fmt"

func main() {
	// n := 2
	const max int = len("hello")
	fmt.Println(max)

	const min = 1
	const pi = 3.14
	const version = "2.0.1"
	const debug = false

	fmt.Println(min, pi, version, debug)
}
