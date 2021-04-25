package main

import "fmt"

func main() {
	i := 100

	switch {
	case i > 100:
		fmt.Println("Big positive number.")
	case i > 0:
		fmt.Println("positive")
	default:
		fmt.Println("number")
	}
}
