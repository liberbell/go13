package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	p := parser{
		sum: make(map[string]int),
	}

	// sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++

		field := strings.Fields(in.Text())
		if len(field) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", field, lines)
			return
		}

		domains := field[0]
		visits, err := strconv.Atoi(field[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (lines #%d)\n", field[1], lines)
		}
	}
}
