package main

import "fmt"

func main() {
	type (
		integer  int
		bookcase [5]int
		cabinet  [5]integer
	)

	blue :=
		// [5]int{6, 9, 3, 2, 4}
		bookcase{6, 9, 3, 2, 4}
	red :=
		// [5]int{6, 9, 3, 2, 1}
		cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal?\n")
	if cabinet(blue) == red {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("Not.")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red:  %#v\n", red)
}
