package main

import "fmt"

type book struct {
	title string
	price float64
}

func printBook(b book) {
	fmt.Printf("%-15s: $%.2f\n", b.tilte, b.price)
}
