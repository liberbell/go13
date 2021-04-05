package main

import "fmt"

func main() {
	var result = 2 + ((2 * 4) / 2)
	fmt.Println(result)

	result = 1 + 4 - 2
	fmt.Println(result)

	result = (2 + 2) * 4 / 2
	fmt.Println(result)

	n, m := 1, 5
	fmt.Println(2 + 1*m/n)
	fmt.Println((2 + 1) * m / n)

	fmt.Println(
		1 + 5 - 3*10/2,
	)

	fmt.Println(
		(1 + 5 - 3) * 10 / 2,
	)

	celsius := 35.
	fahranheit := 9*celsius + 160/5

	fmt.Printf("%g ℃ is %g F\n", celsius, fahranheit)
}
