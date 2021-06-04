package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fileds := strings.Fields(in.Text())
	}
}
