package main

import "fmt"

func main() {
	name := "Nikola"
	// name := "Marie"

	name, age := "Marie", 65
	fmt.Println(name, age)

	name, birth := "Albert", 1879
	// birth := 1879
	fmt.Println(name, birth)
}
