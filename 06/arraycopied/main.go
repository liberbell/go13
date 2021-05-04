package main

import (
	"fmt"
)

func main() {
	// var blue = [3]int{6, 9, 3}

	// red := blue

	// fmt.Println(blue)
	// fmt.Println(red)

	// red[1] = 10
	// fmt.Println(blue)
	// fmt.Println(red)
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	books := prev

	for i := range prev {
		books[i] += " 2nd Edition."
	}
	fmt.Printf("last year: \n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}
