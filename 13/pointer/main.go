package main

import "fmt"

func main() {
	var (
		counter int
		P       *int
	)

	P = &counter

	if P == nil {
		fmt.Printf("P is nil and its address is %p\n", P)
	}

	if P == &counter {
		fmt.Printf("P is equal to counter`s address %p == %p\n", P, &counter)
	}
}
