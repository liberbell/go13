package main

import (
	"fmt"
	"strings"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%q\n", os.Args[i])
	// }

	words := strings.Fields("lazy cat jumps again and again and again")

	var (
		i int
		v string
	)

	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	for i, v = range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
	fmt.Printf("Last value of i and v %d %q\n", i, v)
}
