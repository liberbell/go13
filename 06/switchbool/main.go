package main

import (
	"fmt"
	"os"
)

func main() {
	i := os.Args[1]

	switch {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("Zero")
	}
}
