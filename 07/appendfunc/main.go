package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	var nums []int
	s.Show("no backing array", nums)

	nums = append(nums, 1, 3)
	s.Show("allocate", nums)

	nums = append(nums, 2)
	s.Show("free capacity", nums)

	nums = append(nums, 4)
	s.Show("no allocation", nums)

	nums = append(nums, nums[2:]...)
	s.Show("nums <- 7, 9", nums)

	nums = append(nums[:2], 7, 9)
	s.Show("nums[:2] <- 7, 9", nums)
}
