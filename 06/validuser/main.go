package main

import (
	"fmt"
	"os"
)

const usage = `
usage [username] [password]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}
}
