package main

import "fmt"

func main() {
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	// var running bool
	// running = 1
	// fmt.Println(running)

	// var force float64
	// // speed = force
	// force = 1
	// fmt.Println(force)

	var (
		name   string
		age    int
		famous bool
	)
	name = "Newton"
	age = 84
	famous = true
	fmt.Println(name, age, famous)

	name = "Someone"
	age = 20
	famous = false
	fmt.Println(name, age, famous)

	var prevName string
	prevName = name
	name = "Einstein"
	fmt.Println("Previous name: ", prevName)
	fmt.Println("current name: ", name)

}
