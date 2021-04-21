package main

import (
	"os"
	"strconv"
)

func main() {
	// s := strconv.Itoa(42)
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)
	strconv.Atoi(os.Args[1])
}
