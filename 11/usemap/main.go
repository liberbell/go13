package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		// fmt.Println(in.Text())
		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}
}
