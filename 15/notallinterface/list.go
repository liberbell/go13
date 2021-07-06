package main

import "fmt"

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We`re waiting for delivery."
	}

	for _, p := range l {
		// fmt.Printf("(%-10T) --> ", it)
		// p.print()
		fmt.Printf("* %s\n", p)
	}
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}
