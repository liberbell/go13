package main

import (
	"fmt"
	"strings"
)

func main() {
	start, stop := 'A', 'Z'
	fmt.Println(start, stop)

	fmt.Printf("%s %s\n%s\n", "literal", "dec", strings.Repeat("-", 10))

	for n := start; n <= stop; n++ {
		fmt.Printf("%c -> %[1]d\n", n)
	}
}
