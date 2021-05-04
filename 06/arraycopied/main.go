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

	var books [4]string

	// books := prev

	for i, b := range prev {
		books[i] += b + " 2nd Edition."
	}

	books[3] = "Awesomeness"

	fmt.Printf("last year: \n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}
