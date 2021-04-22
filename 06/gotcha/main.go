package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	if a := os.Args; len(a) != 2 {
		fmt.Println("Input a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", n)
		fmt.Printf("Error: %q\n", err)
	}
}
