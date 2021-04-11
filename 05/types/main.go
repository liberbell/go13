package main

import "fmt"

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

	salt = weight.Gram(100)
	fmt.Printf("Type of weight.Gram: %T\n", weight.Gram(1))
}
