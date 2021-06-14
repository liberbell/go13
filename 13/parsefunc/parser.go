package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse() {
	fields := strings.Fields(in.Text())
	if len(fields) != 2 {
		fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
		return
	}

	domain := fields[0]
	visits, err := strconv.Atoi(fields[1])
	if visits < 0 || err != nil {
		fmt.Printf("wrong input: %q (lines #%d)\n", fields[1], p.lines)
		return
	}
}
