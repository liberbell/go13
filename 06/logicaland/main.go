package main

import "fmt"

func main() {
	// myAge := 25
	// yourAge := 20

	// fmt.Println(myAge && yourAge)
	speed := 100
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	speed = 20
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	speed = 50
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	fmt.Println(1 == 1 && 2 == 2)
}
