package main

import (
	"fmt"
)

func main() {
	var (
		name, lastname string
		age            int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age = "Pablo", "Picasso", 91
	name2, lastname2, age2 = "Sigmund", "Freud", 83

	fmt.Println("Picasso: ", name, lastname, age)
	fmt.Println("Sigmund: ", name2, lastname2, age2)

	type person struct {
		name, lastname string
		age            int
	}

	// var picasso person
	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}

	var freud person

	// picasso.name = "Pablo"
	// picasso.lastname = "Picasso"
	// picasso.age = 91

	freud.name = "Sigmund"
	freud.lastname = "Freud"

	fmt.Printf("\n%s age is %d\n", picasso.lastname, picasso.age)

	fmt.Printf("\nPicasso: %#v\n", picasso)
	fmt.Printf("Freud: %#v\n", freud)
}
