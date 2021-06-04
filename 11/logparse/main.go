package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		domain := fields[0]
		visits := fields[1]
		fmt.Printf("domain: %s - visits: %s\n", domain, visits)
	}
}
