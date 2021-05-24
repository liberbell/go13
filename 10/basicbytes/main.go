package main

import "fmt"

func main() {
	start, stop := 'A', 'Z'
	fmt.Println(start, stop)

	fmt.Printf("%s\n%s\n", "literal")

	for n := start; n <= stop; n++ {
		fmt.Printf("%c -> %[1]\nd ", n)
	}
}
