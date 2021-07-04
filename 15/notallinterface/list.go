package main

import "fmt"

type list []*product

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We`re waiting for delivery.")
		return
	}

	for _, p := range l {
		// fmt.Printf("(%-10T) --> ", it)
		p.print()
	}
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}
