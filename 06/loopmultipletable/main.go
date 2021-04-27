package main

import "fmt"

func main() {
	fmt.Printf("%5s", "X")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%5d", i)
	}
}
