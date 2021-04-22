package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	if a := os.Args; len(a) != 2 {
		fmt.Println("Input a number.")
	}
}
