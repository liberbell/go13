package main

import (
	"fmt"
	"math"
)

func main() {
	// var (
	// 	width  uint8 = 255
	// 	height       = 255
	// )

	// width++
	// if int(width) < height {
	// 	// fmt.Println(width)
	// 	fmt.Println("height is greater.")
	// }
	// fmt.Printf("width: %d, height: %d\n", width, height)
	// int8(math.MaxInt8 + 1)
	var n int8 = math.MaxInt8
	fmt.Println("max int8    :", n)
	fmt.Println("max int8 + 1:", n+1)

	n = math.MinInt8
	fmt.Println("min int8    :", n)
	fmt.Println("min int8 -1 :", n-1)

	var un uint8
	fmt.Println("min uint8    :", un)
	fmt.Println("min uint8 +1 :", un+1)
}
