package main

import "fmt"

func main() {
	var P *int

	if P == nil {
		fmt.Printf("P is nil and its address is %p\n", P)
	}
}
