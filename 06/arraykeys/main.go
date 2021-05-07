package main

import "fmt"

const (
	ETH = iota
	WAN
)

func main() {
	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}
	fmt.Printf("1 BTS is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTS is %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
