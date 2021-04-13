package main

import "fmt"

func main() {
	// const min = 1 + 1
	// const pi = 3.14 * min

	// fmt.Printf("min is %d, pi is %g\n", min, pi)

	// const min = 42
	// // var i int
	// var f float64
	// f = min

	// fmt.Println(f)
	const min = 42
	var i int = min
	var f float64 = min
	var b byte = min
	var j int32 = min

	fmt.Println(min, i, f, b, j)
	fmt.Printf("Type of i: %T, Type of f: %T, Type of b: %T, Type of j: %T", i, f, b, j)
}
