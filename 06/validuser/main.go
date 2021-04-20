package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage     = "usage: [username] [password]"
	errUser   = "Access Denied for %q\n"
	errPass   = "Invalid password for %q.\n"
	accGrant  = "Access granted for %q\n"
	username1 = "jack"
	password1 = "p@ss"
	username2 = "bety"
	password2 = "puss"
)

func main() {
	arg := os.Args

	if len(arg) != 3 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	user, pass := os.Args[1], os.Args[2]
	// if user != username1 {
	// 	fmt.Printf(errUser, user)
	// } else if pass == password1 {
	// 	fmt.Println(accGrant)
	// } else {
	// 	fmt.Printf(errPass, user)
	// }
	if user == username1 && pass == password1 {
		fmt.Printf(accGrant, user)
	} else if user == username2 && pass == password2 {
		fmt.Printf(accGrant, user)
	} else if user == username1 && pass != password1 {
		fmt.Printf(errPass, user)
	} else if user == username2 && pass != password2 {
		fmt.Printf(errPass, user)
	} else if user != username1 || user != username2 {
		fmt.Printf(errUser, user)
	} else {
		fmt.Println(usage)
	}
}
