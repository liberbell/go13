package main

import "github.com/inancgumus/prettyslice"

func main() {
	evens := []int{2, 4}
	odds := []int{1, 3, 5}
	numbers := copy(evens, odds)

	prettyslice.Show("evens", evens)
	prettyslice.Show("odds", odds)

	prettyslice.Show("numbers", numbers)
}
