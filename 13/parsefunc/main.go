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
}
