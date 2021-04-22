package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [number]")
		return
	}
	arg := os.Args[1]
	number, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println("Error: Input number.")
		return
	}
	fmt.Printf("%g * 2 is %g.\n", number, number*2)
}
