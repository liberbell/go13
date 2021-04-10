package main

import "fmt"

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce

	o = g * 0.035274
	fmt.Println(o)
}
