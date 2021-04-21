package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// s := strconv.Itoa(42)
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println(n, err)
}
