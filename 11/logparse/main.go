package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
	)
	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		domain := fields[0]
		visits, _ := strconv.Atoi(fields[1])

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		sum[domain] += visits
	}
	fmt.Printf("%-30s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 45))

	for domain, visits := range sum {
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
}
