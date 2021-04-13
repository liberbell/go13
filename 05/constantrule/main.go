package main

import "fmt"

func main() {
	// n := 2
	const max int = len("hello")
	fmt.Println(max)

	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = false

	fmt.Println(min, pi, version, debug)
}
