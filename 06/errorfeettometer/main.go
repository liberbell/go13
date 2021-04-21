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
}
