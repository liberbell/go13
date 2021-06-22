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
	lerr    error
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p *parser, line string) (parsed result, err error) {
	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error

	parsed.domain = fields[0]

	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (lines #%d)", fields[1], p.lines)
	}
	return
}

func update(p *parser, parsed result) {
	fmt.Printf("update.p     : %p - %p\n\n", p, &p)
	domain, visits := parsed.domain, parsed.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}
	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		// visits: visits + p.sum[domain].visits,
	}
}
