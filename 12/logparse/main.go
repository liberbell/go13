package main

import (
	"bufio"
	"os"
)

type parser struct {
	sum map[string]int
}

func main() {
	var (
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
}
