package main

import (
	"os"
	"strconv"
)

func main() {
	arg := os.Args[2]

	celsius, _ := strconv.ParseFloat(arg, float64)
}
