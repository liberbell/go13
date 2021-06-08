package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		domain := field[0]

		visits, err := strconv.Atoi(field[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (lines #%d)\n", field[1], lines)
			return
		}

		if _, ok := p.sum[domain]; ok {
			domains = append(domains, domain)
		}
		total += visits
		p.sum[domain] += visits
	}
	sort.Strings(domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range domains {
		visits := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
}
