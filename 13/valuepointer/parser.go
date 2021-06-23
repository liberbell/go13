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

func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

func parse(p *parser, line string) (r result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error

	r.domain = fields[0]

	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (lines #%d)", fields[1], p.lines)
	}
	return
}

func update(p *parser, r result) {
	fmt.Printf("update.p     : %p - %p\n\n", p, &p)
	domain, visits := r.domain, r.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}
	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		// visits: visits + p.sum[domain].visits,
	}
}

func err(p parser) error {
	return p.lerr
}
