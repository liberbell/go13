package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := nums[:2:2]
	odds = append(odds, 5, 7)

	s.Show("nums", nums)
	s.Show("odds", odds)
}
