package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We`re waiting for delivery.")
		return
	}

	for _, it := range l {
		// fmt.Printf("(%-10T) --> ", it)
		it.print()
	}
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		g, ok := it.(interface{ discount(ratio) })
		if !ok {
			continue
		}
		g.discount(ratio)
	}
}
