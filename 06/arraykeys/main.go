package main

import "fmt"

func main() {
	rates := [...]float64{
		25.5,
		120.5,
	}
	fmt.Printf("1 BTS is %g ETH\n", rates[0])
}
