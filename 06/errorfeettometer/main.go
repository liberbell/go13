package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
Feet to meters
--------------
This program converts feet into meters.
usage:
feet[feetsToConverts]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]
	feet, _ := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", arg, feet)
}
