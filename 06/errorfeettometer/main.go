package main

import (
	"fmt"
	"os"
)

const usage = `
Feet to meters
--------------
This program converts feet into meters.
usage:
feet[feetsToConverts]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSapce(usage))
		return
	}
}
