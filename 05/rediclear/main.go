package main

import "fmt"

func main() {
	name := "Nikola"
	// name := "Marie"

	name, age := "Marie", 65
	fmt.Println(name, age)

	name = "Albert"
	birth := 1879
	fmt.Println(name, birth)
}
