package main

import (
	"bufio"
	"os"
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
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
	}
}
