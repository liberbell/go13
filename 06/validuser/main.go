package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = "usage: [username] [password]"
	errUser = "Access Denied for %q\n"

func main() {
	arg := os.Args

	if len(arg) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	user, pass := os.Args[1], os.Args[2]
	if user != "jack" {
		fmt.Printf(errUser, user)
	} else if pass == "p@ss" {
		fmt.Println("Access granted.")
	} else {
		fmt.Printf("Invalid password for %q.\n", user)
	}
}
