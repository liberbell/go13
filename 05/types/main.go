package main

import (
	"fmt"

	"./weights"
)

func main() {
	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt  Gram     = 100
		apple Kilogram = 5
		truck Ton      = 10
	)

	fmt.Printf("salt: %d, apple: %d, truck: %d\n", salt, apple, truck)

	salt = Gram(apple)
	apple = Kilogram(truck)
	truck = Ton(Gram(int(apple)))

	salt = Gram(weights.Gram(100))
	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
}
