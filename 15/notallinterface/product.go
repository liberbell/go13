package main

import (
	"fmt"
)

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)\n", p.title, p.price, p.released.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
