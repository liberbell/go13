package main

import s "github.com/inancgumus/prettyslice"

func main() {
	agesArray := [3]int{35, 15, 25}
	// age := agesArray[0:3]

	// fmt.Println(agesArray)
	// fmt.Println(age)

	// age[0] = 100
	// fmt.Println(agesArray)
	// fmt.Println(age)

	s.Show(agesArray)

}
