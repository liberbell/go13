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
		// fmt.Println("please tell me a value in feet")
		fmt.Println(strings.TrimSpace(usage))
		return
	}
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048

	fmt.Printf("%g feet(s) is %g meters.\n", feet, meters)
}
