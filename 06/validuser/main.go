package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage    = "usage: [username] [password]"
	errUser  = "Access Denied for %q\n"
	errPass  = "Invalid password for %q.\n"
	accGrant = "Access granted."
	username = "jack"
	password = "p@ss"
)

func main() {
	arg := os.Args

	if len(arg) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	user, pass := os.Args[1], os.Args[2]
	if user != username {
		fmt.Printf(errUser, user)
	} else if pass == password {
		fmt.Println(accGrant)
	} else {
		fmt.Printf(errPass, user)
	}
}
