package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum map[string]int
	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		domain := fields[0]
		visits, _ := strconv.Atoi(fields[1])

		sum[domain] += visits
	}
	fmt.Printf("%-30s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 45))
}
