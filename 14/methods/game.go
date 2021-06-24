package main

import "fmt"

type game struct {
	title string
	price float64
}

func printGame(g game) {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
