package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	i := int(input)

	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("Zero")
	}
}
