package main

import "fmt"

func main() {
	const min = 1
	max := min + 5

	fmt.Printf("Min: %d and Type is %T, Max: %d Type is %T\n", min, min, max, max)

	i := 42 + 3.14
	f := 3.14
	b := true
	s := "hello"
	r := `A`

	fmt.Printf("Type i: %T, Type f: %T, Type b: %T, Type s: %T, Type r: %T\n", i, f, b, s, r)
}
