package main

import "fmt"

func main() {
	blue :=
		[5]int{6, 9, 3, 2, 1}
	red :=
		[5]int{6, 9, 3, 2, 1}

	fmt.Print("Are they equal?\n")
	if blue == red {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("Not.")
	}
}
