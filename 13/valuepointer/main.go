package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := parse(p, in.Text())
		update(p, parsed)
	}
	summarize(p)
	dumpErr([]error{in.Err(), err(p)})
}

func dumpErr(errs []error) {
	// if p.lerr != nil {
	// 	fmt.Println("> Err:", p.lerr)
	// }

	// if err := in.Err(); err != nil {
	// 	fmt.Println("> Err:", err)
	// }

	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}

func summarize(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		// parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}
