package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
		fmt.Println(q)
	}
}
