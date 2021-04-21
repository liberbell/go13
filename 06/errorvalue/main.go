package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
