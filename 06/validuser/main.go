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

	if len(arg) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	user := os.Args[1]
	if user != "jack" {
		fmt.Printf("Access Denied for %q\n", user)
	}

	pass := os.Args[2]
	if pass != "p@ss" {
		fmt.Println("Invalid password.")
	}
}
