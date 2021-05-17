package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	s.Show("nums", nums)
}
