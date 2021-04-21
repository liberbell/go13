package main

import (
	"fmt"
	"os"
	"strconv"
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
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: %q\n", err)
	}
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
