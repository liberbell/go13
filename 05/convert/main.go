package main

import (
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	strconv.ParseFloat(arg, 64)
}
