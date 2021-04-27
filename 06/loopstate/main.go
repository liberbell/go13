package main

import "fmt"

func main() {
	var sum int

	// sum++
	// sum += 2
	// sum += 3
	// sum += 4
	// sum += 5

	// fmt.Println(sum)

	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
