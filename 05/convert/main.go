package main

import (
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)
}
