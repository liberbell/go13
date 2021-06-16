package main

import (
	"bufio"
	"os"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		parsed, err := parse(p, in.Text())
		domain, visits := parsed.domain, parsed.visits

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		p.total += visits

		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}
}
