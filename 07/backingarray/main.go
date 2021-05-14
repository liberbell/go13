package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// agesArray := [3]int{35, 15, 25}
	// age := agesArray[0:3]

	// fmt.Println(agesArray)
	// fmt.Println(age)

	// age[0] = 100
	// fmt.Println(agesArray)
	// fmt.Println(age)

	grades := [...]float64{40, 10, 20, 50, 60, 70}
	front := grades[:3]

	s.PrintBacking = true
	s.MaxPerLine = 7

	s.Show("grade", grades[:])
	s.Show("front", front)

}
