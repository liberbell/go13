package main

import "fmt"

func main() {
	width, height := 5., 12.

	area := width * height
	fmt.Printf("Area is : %g\n", area)

	area = area - 10
	fmt.Printf("Area is : %g\n", area)

	area = area + 10
	fmt.Printf("Area is : %g\n", area)
}
