package main

import (
	"fmt"
)

func main() {
	var blue = [3]int{6, 9, 3}

	red := blue

	fmt.Println(blue)
	fmt.Println(red)

	red[1] = 10
	fmt.Println(blue)
	fmt.Println(red)
}
