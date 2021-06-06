package main

import "fmt"

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

	var picasso struct {
		name     string
		lastname string
		age      int
	}

	var freud struct {
		name, lastname string
		age            int
	}
}
