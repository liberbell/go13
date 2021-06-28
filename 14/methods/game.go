package main

import "fmt"

type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Sprintf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}
