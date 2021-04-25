package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// i := -10
	input := os.Args[1]
	i, _ := strconv.Atoi(input)

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
