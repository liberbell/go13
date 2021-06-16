package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p parser, line string) {
	var (
		parsed result
		err    error
	)

	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
		return parsed
	}

	parsed.domain := fields[0]

	parsed.visits, err := strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Printf("wrong input: %q (lines #%d)\n", fields[1], p.lines)
		return parsed
	}
	return parsed
}
