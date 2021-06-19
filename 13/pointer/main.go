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

	counter = 10
	fmt.Printf("counter: %-13d addr: %13p\n", counter, &counter)

	fmt.Println("\n......... change counter in passVal()")
	passVal(counter)
	fmt.Printf("counter: %-13d addr: %13p\n", counter, &counter)

	fmt.Println("\n......... change counter in passPtrVal()")
	passPtrVal(&counter)
	fmt.Printf("counter: %-13d addr: %13p\n", counter, &counter)
}

func passPtrVal(pn *int) {
	fmt.Printf("pn      : %-13d addr: %13p *pn: %-13d\n", pn, &pn, *pn)
	(*pn)++
}

func passVal(n int) {
	n = 50
	fmt.Printf("n      : %-13d addr: %13p\n", n, &n)
}
