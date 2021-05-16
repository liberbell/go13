package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	var nums []int
	s.Show("no backing array", nums)
}
