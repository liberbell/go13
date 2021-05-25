package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// start, stop := 'A', 'Z'
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	fmt.Println(start, stop)

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x  % -12x\n", n, string(n))
	}
}
