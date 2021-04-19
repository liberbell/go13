package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
usage [username] [password]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}
}
