package main

import "fmt"

func main() {
	var (
		counter int
		V       int
		P       *int
	)

	counter = 100

	P = &counter
	V = *P

	// if P == nil {
	// 	fmt.Printf("P is nil and its address is %p\n", P)
	// }

	// if P == &counter {
	// 	fmt.Printf("P is equal to counter`s address %p == %p\n", P, &counter)
	// }

	fmt.Printf("counter: %-13d addr: %13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %13p\n", V, &V)

	fmt.Println("\n......... change counter")
}
