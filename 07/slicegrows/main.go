package main

import (
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	// s.PrintBacking = true
	// s.MaxPerLine = 30
	// s.Width = 150

	// var nums []int

	// screen.Clear()
	// for cap(nums) <= 128 {
	// 	screen.MoveTopLeft()

	// 	s.Show("nums", nums)
	// 	nums = append(nums, rand.Intn(9)+1)

	// 	time.Sleep(time.Second / 4)
	// }

	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := float64(cap(ages))
		if c := oldCap
	}
}
