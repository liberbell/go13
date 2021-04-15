package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const (
		feetInMeters = 0.3048
		feetInYards  = math.Round(feetInMeters) / 0.9144
	)
	fmt.Printf("%T\n", feetInYards)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := feet * feetInYards

	fmt.Printf("%g feet is %g meters\n", feet, meters)
	fmt.Printf("%g feet is %g yards\n", feet, yards)
}
