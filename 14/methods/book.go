package main

import "fmt"

type book struct {
	title string
	price money
}

func (b book) print() {
	fmt.Sprintf("%-15s: %s\n", b.title, b.price.string())
}
