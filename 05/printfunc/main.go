package main

import "fmt"

func main() {
	var brand = "Google"
	fmt.Printf("%s\n", brand)

	ops := 250
	ok := 543
	fail := 433

	fmt.Printf("total: %d success: %d / %d\n", ops, ok, fail)
}
