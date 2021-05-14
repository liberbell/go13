package main

import (
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// agesArray := [3]int{35, 15, 25}
	// age := agesArray[0:3]

	// fmt.Println(agesArray)
	// fmt.Println(age)

	// age[0] = 100
	// fmt.Println(agesArray)
	// fmt.Println(age)

	// grades := [...]float64{40, 10, 20, 50, 60, 70}
	grades := []float64{40, 10, 20, 50, 60, 70}

	var newGrades []float64
	newGrades = append(newGrades, grades...)
	front := newGrades[:3]
	front2 := front[:3]
	front3 := front

	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7

	s.Show("grade", grades[:])
	s.Show("newGrade", newGrades)
	s.Show("front", front)
	s.Show("front2", front2)
	s.Show("front3", front3)

}
