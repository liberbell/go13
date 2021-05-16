package main

import (
	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
	}
}
