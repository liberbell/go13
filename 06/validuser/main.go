package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
usage: [username] [password]`

func main() {
	arg := os.Args

	if arg < 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}
}
