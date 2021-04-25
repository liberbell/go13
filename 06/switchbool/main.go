package main

import (
	"fmt"
	"os"
)

func main() {

	switch {
		i := os.Args[1]
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("Zero")
	}
}
