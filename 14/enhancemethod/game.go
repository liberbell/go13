package main

import "fmt"

type game struct {
	title string
	price float64
}

func printBook(g game) {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
