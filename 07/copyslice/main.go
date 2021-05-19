package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	// evens := []int{2, 4}
	// odds := []int{1, 3, 5}

	// s.Show("evens", evens)
	// s.Show("odds", odds)

	// numbers := copy(evens, odds)

	// s.Show("numbers", numbers)

	data := []float64{10, 25, 30, 50}
	// newData := []float64{99, 100}
	// for i := range newData {
	// 	data[i] = newData[i]
	// }
	copy(data, []float64{99, 100})
	copy(data, []float64{10, 5, 15, 0, 20})

	s.Show("Probabilies", data)
	fmt.Printf("It is gonna rain? %.f%% chance.\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))
}
