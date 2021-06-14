package main

import "bufio"

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

	in := bufio.NewScanner(os.stdin)
}
