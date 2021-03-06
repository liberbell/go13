package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: [number]")
	// }
	// arg := os.Args[1]
	// number, err := strconv.ParseFloat(arg, 64)
	// if err != nil {
	// 	fmt.Println("Error: Input number.")
	// 	return
	// }
	// fmt.Printf("%g * 2 is %g.\n", number, number*2)

	if arg := os.Args; len(arg) != 2 {
		fmt.Println("Error: Input a number.")
	} else if num, err := strconv.Atoi(arg[1]); err != nil {
		fmt.Printf("Cannot convert: %q\n", arg[1])
	} else {
		fmt.Printf("%s * 2 is %d\n", arg[1], num*2)
	}

}
