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
		fmt.Printf("Cannot convert %q.\n", a[1])
		fmt.Printf("Error: %q\n", err)
	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}
	fmt.Printf("n is %d. 笑 笑 笑 - \n", n)
}
