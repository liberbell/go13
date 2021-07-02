package main

import "fmt"

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	return "unknown"
}
