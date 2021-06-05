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
		total   int
	)
	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v\n", fields)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("Wrong input: %q\n", fields[1])
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		total += visits
		sum[domain] += visits
	}
	fmt.Printf("%-30s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 45))

	// for domain, visits := range sum {
	// 	fmt.Printf("%-30s %10d\n", domain, visits)
	// }
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", total)
}
