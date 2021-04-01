package main

import "fmt"

func main() {
	// speed := 100
	// force := 2.5

	// // speed = speed * int(force)
	// speed = int(float64(speed) * force)
	// fmt.Println(speed)

	var apple int
	var orange int32

	apple = int(orange)
	orange = 65
	color := string(orange)

	fmt.Println(color)
	fmt.Println(string(65.0))

	fmt.Println(apple)
}
