package main

import "fmt"

func main() {
	// var brand = "Google"
	// fmt.Printf("%s\n", brand)

	// ops := 250
	// ok := 543
	// fail := 433

	// fmt.Printf("total: %d success: %d / %d\n", ops, ok, fail)
	// var speed int
	// var heat float64
	// var off bool
	// var brand string

	// fmt.Printf("%T\n", speed)
	// fmt.Printf("%T\n", heat)
	// fmt.Printf("%T\n", off)
	// fmt.Printf("%T\n", brand)
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.611
		hasLife  = false
	)
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v has Life?: %v\n", planet, hasLife)

	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)
}
