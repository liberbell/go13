package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	p := parser{
		sum: make(map[string]result),
	}

	// sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		field := strings.Fields(in.Text())
		if len(field) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", field, p.lines)
			return
		}

		domain := field[0]

		visits, err := strconv.Atoi(field[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (lines #%d)\n", field[1], p.lines)
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		p.total += visits
		r := result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
		// p.sum[domain].result += visits
		p.sum[domain] = r
	}
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	fmt.Printf("\n%-30s %10d\n", "DOMAIN", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err: ", err)
	}
}
