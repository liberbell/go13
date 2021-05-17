package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := nums[:2]
	s.Show("nums", nums)
	s.Show("odds", odds)
}
