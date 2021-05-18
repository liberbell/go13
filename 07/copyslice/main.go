package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	evens := []int{2, 4}
	odds := []int{1, 3, 5}

	s.Show("evens", evens)
	s.Show("odds", odds)

	numbers := copy(evens, odds)

	s.Show("numbers", numbers)
}
