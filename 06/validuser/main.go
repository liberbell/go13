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

	user, pass := os.Args[1], os.Args[2]
	if user != "jack" {
		fmt.Printf("Access Denied for %q\n", user)
	} else if pass != "p@ss" {
		fmt.Println("Invalid password.")
	}
}
