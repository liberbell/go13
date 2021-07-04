package main

import (
	"fmt"
)

type book struct {
	product
}

func (b *book) print() {
	b.product.print()

	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}
