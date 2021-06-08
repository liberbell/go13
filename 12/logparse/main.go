package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type parser struct {
	sum map[string]int
}

func main() {
	var (
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++

		field := strings.Fields(in.Text())
		if len(field) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", field, lines)
		}
	}
}
