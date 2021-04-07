package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	celsius, _ := strconv.ParseFloat(arg, 64)
	fahrenheit := (celsius*9 + 5) + 32
	fmt.Printf("%g °C is %g °F\n", celsius, fahrenheit)
}
