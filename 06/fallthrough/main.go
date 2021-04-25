package main

import "fmt"

func main() {
	i := -10

	switch {
	case i > 100:
		fmt.Println("Big positive number.")
		fallthrough
	case i > 0:
		fmt.Println("positive")
		fallthrough
	default:
		fmt.Println("number")
	}
}
