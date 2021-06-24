package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) printGame() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
