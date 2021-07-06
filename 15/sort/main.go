package main

import (
	"fmt"
	"sort"
)

func main() {
	l := list{
		{title: "mody dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 10, released: toTimestamp("723622855")},
		{title: "hobbit", price: 25},
	}

	sort.Sort(l)
	fmt.Print(l)
}
