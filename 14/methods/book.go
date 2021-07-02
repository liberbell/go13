package main

import "fmt"

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	if v == nil {
		return "unknown"
	}
	return ""
}
